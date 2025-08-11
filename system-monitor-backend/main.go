package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/disk"
)

// corsMiddleware adds CORS headers to all responses.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func getCPUUsage(w http.ResponseWriter, r *http.Request) {
	percent, err := cpu.Percent(time.Second, false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]float64{"cpu_usage": percent[0]})
}

func getMemoryUsage(w http.ResponseWriter, r *http.Request) {
	v, err := mem.VirtualMemory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]uint64{"total": v.Total, "available": v.Available, "used": v.Used, "free": v.Free, "used_percent": uint64(v.UsedPercent)})
}

func getNetworkUsage(w http.ResponseWriter, r *http.Request) {
	netIO, err := net.IOCounters(false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if len(netIO) > 0 {
		json.NewEncoder(w).Encode(map[string]uint64{"bytes_sent": netIO[0].BytesSent, "bytes_recv": netIO[0].BytesRecv})
	} else {
		json.NewEncoder(w).Encode(map[string]string{"message": "No network interfaces found"})
	}
}

func getDiskUsage(w http.ResponseWriter, r *http.Request) {
	parts, err := disk.Partitions(true)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var diskInfo []map[string]interface{}
	for _, p := range parts {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			log.Printf("Error getting disk usage for %s: %v", p.Mountpoint, err)
			continue
		}
		diskInfo = append(diskInfo, map[string]interface{}{
			"path": p.Mountpoint,
			"total": usage.Total,
			"free": usage.Free,
			"used": usage.Used,
			"used_percent": usage.UsedPercent,
		})
	}
	json.NewEncoder(w).Encode(diskInfo)
}

func main() {
	http.Handle("/api/cpu", corsMiddleware(http.HandlerFunc(getCPUUsage)))
	http.Handle("/api/memory", corsMiddleware(http.HandlerFunc(getMemoryUsage)))
	http.Handle("/api/network", corsMiddleware(http.HandlerFunc(getNetworkUsage)))
	http.Handle("/api/disk", corsMiddleware(http.HandlerFunc(getDiskUsage)))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterOutlet } from '@angular/router';
import { SystemMonitorService } from './system-monitor.service';

@Component({
  selector: 'app-root',
  standalone: true,
  imports: [CommonModule, RouterOutlet],
  templateUrl: './app.component.html',
  styleUrl: './app.component.css'
})
export class AppComponent implements OnInit {
  title = 'system-monitor-frontend';
  cpuUsage: any;
  memoryUsage: any;
  networkUsage: any;
  diskUsage: any;

  constructor(private systemMonitorService: SystemMonitorService) { }

  ngOnInit(): void {
    this.getSystemMetrics();
    setInterval(() => this.getSystemMetrics(), 5000); // Actualizar cada 5 segundos
  }

  getSystemMetrics(): void {
    this.systemMonitorService.getCPUUsage().subscribe(data => {
      this.cpuUsage = data.cpu_usage;
    });

    this.systemMonitorService.getMemoryUsage().subscribe(data => {
      this.memoryUsage = data;
    });

    this.systemMonitorService.getNetworkUsage().subscribe(data => {
      this.networkUsage = data;
    });

    this.systemMonitorService.getDiskUsage().subscribe(data => {
      this.diskUsage = data;
    });
  }
}
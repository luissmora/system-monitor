import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SystemMonitorService {
  private apiUrl = 'http://localhost:8080/api';

  constructor(private http: HttpClient) { }

  getCPUUsage(): Observable<any> {
    return this.http.get(`${this.apiUrl}/cpu`);
  }

  getMemoryUsage(): Observable<any> {
    return this.http.get(`${this.apiUrl}/memory`);
  }

  getNetworkUsage(): Observable<any> {
    return this.http.get(`${this.apiUrl}/network`);
  }

  getDiskUsage(): Observable<any> {
    return this.http.get(`${this.apiUrl}/disk`);
  }
}
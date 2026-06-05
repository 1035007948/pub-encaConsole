import { API_ENDPOINTS } from './config';

export interface ApiResponse<T> {
  total?: number;
  items?: T[];
  error?: string;
}

export interface Complaint {
  id: number;
  complaint_no: string;
  title: string;
  description: string;
  status: string;
  level: string;
  complainant_name: string;
  complainant_tel: string;
  enterprise_name: string;
  enterprise_addr: string;
  responsible_user: string;
  batch_no: string;
  priority: number;
  remark: string;
  created_at: string;
  updated_at: string;
}

export interface SamplingPoint {
  id: number;
  point_no: string;
  point_name: string;
  address: string;
  longitude: number;
  latitude: number;
  status: string;
  complaint_id: number;
  complaint_no: string;
  responsible_user: string;
  scheduled_date: string | null;
  scheduled_time_from: string;
  scheduled_time_to: string;
  batch_no: string;
  remark: string;
  created_at: string;
  updated_at: string;
}

export interface NoiseReading {
  id: number;
  reading_no: string;
  sampling_point_id: number;
  point_no: string;
  complaint_id: number;
  complaint_no: string;
  time_period_id: number;
  period_name: string;
  measurement_date: string;
  measurement_time: string;
  leq: number;
  lmax: number;
  lmin: number;
  l10: number;
  l90: number;
  standard_limit: number;
  exceed_value: number;
  is_exceeded: boolean;
  status: string;
  responsible_user: string;
  batch_no: string;
  remark: string;
  created_at: string;
  updated_at: string;
}

export interface TimePeriod {
  id: number;
  period_no: string;
  period_name: string;
  period_type: string;
  time_from: string;
  time_to: string;
  day_limit: number;
  night_limit: number;
  status: string;
  description: string;
  batch_no: string;
  remark: string;
  created_at: string;
  updated_at: string;
}

export interface AnomalyEvent {
  id: number;
  event_no: string;
  event_name: string;
  event_type: string;
  severity: string;
  entity_type: string;
  entity_id: number;
  entity_no: string;
  trigger_field: string;
  trigger_value: string;
  threshold_value: string;
  description: string;
  status: string;
  responsible_user: string;
  deadline: string | null;
  resolved_at: string | null;
  resolution_note: string;
  rule_id: number;
  rule_no: string;
  batch_no: string;
  remark: string;
  created_at: string;
  updated_at: string;
}

export interface StatisticsDashboard {
  total_complaints: number;
  pending_complaints: number;
  completed_complaints: number;
  total_sampling_points: number;
  total_noise_readings: number;
  exceeded_readings: number;
  average_leq: number;
  evidence_completeness: number;
  rectification_rate: number;
  retest_pass_rate: number;
  open_anomalies: number;
}

async function fetchApi<T>(url: string, options?: RequestInit): Promise<T> {
  const response = await fetch(url, {
    ...options,
    headers: {
      'Content-Type': 'application/json',
      ...options?.headers,
    },
  });

  if (!response.ok) {
    const error = await response.json();
    throw new Error(error.error || 'API request failed');
  }

  return response.json();
}

export const complaintsApi = {
  list: (params?: Record<string, string>) => {
    const query = params ? `?${new URLSearchParams(params)}` : '';
    return fetchApi<ApiResponse<Complaint>>(`${API_ENDPOINTS.complaints.list}${query}`);
  },
  get: (id: number) => fetchApi<Complaint>(API_ENDPOINTS.complaints.get(id)),
  create: (data: Partial<Complaint>) =>
    fetchApi<Complaint>(API_ENDPOINTS.complaints.create, {
      method: 'POST',
      body: JSON.stringify(data),
    }),
  update: (id: number, data: Partial<Complaint>) =>
    fetchApi<Complaint>(API_ENDPOINTS.complaints.update(id), {
      method: 'PUT',
      body: JSON.stringify(data),
    }),
  delete: (id: number) =>
    fetchApi<{ message: string }>(API_ENDPOINTS.complaints.delete(id), {
      method: 'DELETE',
    }),
  transition: (id: number, toStatus: string, reason?: string) =>
    fetchApi<Complaint>(API_ENDPOINTS.complaints.transition(id), {
      method: 'POST',
      body: JSON.stringify({ to_status: toStatus, reason }),
    }),
};

export const samplingPointsApi = {
  list: (params?: Record<string, string>) => {
    const query = params ? `?${new URLSearchParams(params)}` : '';
    return fetchApi<ApiResponse<SamplingPoint>>(`${API_ENDPOINTS.samplingPoints.list}${query}`);
  },
  get: (id: number) => fetchApi<SamplingPoint>(API_ENDPOINTS.samplingPoints.get(id)),
  create: (data: Partial<SamplingPoint>) =>
    fetchApi<SamplingPoint>(API_ENDPOINTS.samplingPoints.create, {
      method: 'POST',
      body: JSON.stringify(data),
    }),
  update: (id: number, data: Partial<SamplingPoint>) =>
    fetchApi<SamplingPoint>(API_ENDPOINTS.samplingPoints.update(id), {
      method: 'PUT',
      body: JSON.stringify(data),
    }),
  delete: (id: number) =>
    fetchApi<{ message: string }>(API_ENDPOINTS.samplingPoints.delete(id), {
      method: 'DELETE',
    }),
  transition: (id: number, toStatus: string, reason?: string) =>
    fetchApi<SamplingPoint>(API_ENDPOINTS.samplingPoints.transition(id), {
      method: 'POST',
      body: JSON.stringify({ to_status: toStatus, reason }),
    }),
  validate: (id: number) =>
    fetchApi<{ is_valid: boolean; missing_fields: string[] }>(
      API_ENDPOINTS.samplingPoints.validate(id)
    ),
};

export const noiseReadingsApi = {
  list: (params?: Record<string, string>) => {
    const query = params ? `?${new URLSearchParams(params)}` : '';
    return fetchApi<ApiResponse<NoiseReading>>(`${API_ENDPOINTS.noiseReadings.list}${query}`);
  },
  get: (id: number) => fetchApi<NoiseReading>(API_ENDPOINTS.noiseReadings.get(id)),
  create: (data: Partial<NoiseReading>) =>
    fetchApi<NoiseReading>(API_ENDPOINTS.noiseReadings.create, {
      method: 'POST',
      body: JSON.stringify(data),
    }),
  update: (id: number, data: Partial<NoiseReading>) =>
    fetchApi<NoiseReading>(API_ENDPOINTS.noiseReadings.update(id), {
      method: 'PUT',
      body: JSON.stringify(data),
    }),
  delete: (id: number) =>
    fetchApi<{ message: string }>(API_ENDPOINTS.noiseReadings.delete(id), {
      method: 'DELETE',
    }),
  byPoint: (id: number) =>
    fetchApi<ApiResponse<NoiseReading>>(API_ENDPOINTS.noiseReadings.byPoint(id)),
};

export const timePeriodsApi = {
  list: (params?: Record<string, string>) => {
    const query = params ? `?${new URLSearchParams(params)}` : '';
    return fetchApi<ApiResponse<TimePeriod>>(`${API_ENDPOINTS.timePeriods.list}${query}`);
  },
  get: (id: number) => fetchApi<TimePeriod>(API_ENDPOINTS.timePeriods.get(id)),
  create: (data: Partial<TimePeriod>) =>
    fetchApi<TimePeriod>(API_ENDPOINTS.timePeriods.create, {
      method: 'POST',
      body: JSON.stringify(data),
    }),
  update: (id: number, data: Partial<TimePeriod>) =>
    fetchApi<TimePeriod>(API_ENDPOINTS.timePeriods.update(id), {
      method: 'PUT',
      body: JSON.stringify(data),
    }),
  delete: (id: number) =>
    fetchApi<{ message: string }>(API_ENDPOINTS.timePeriods.delete(id), {
      method: 'DELETE',
    }),
  batchImport: (items: Partial<TimePeriod>[]) =>
    fetchApi<{ total: number; success: number; failed: number; errors: string[]; warnings: string[] }>(
      API_ENDPOINTS.timePeriods.batchImport,
      {
        method: 'POST',
        body: JSON.stringify({ items }),
      }
    ),
};

export const anomaliesApi = {
  list: (params?: Record<string, string>) => {
    const query = params ? `?${new URLSearchParams(params)}` : '';
    return fetchApi<ApiResponse<AnomalyEvent>>(`${API_ENDPOINTS.anomalies.list}${query}`);
  },
  get: (id: number) => fetchApi<AnomalyEvent>(API_ENDPOINTS.anomalies.get(id)),
  resolve: (id: number, resolutionNote: string) =>
    fetchApi<{ message: string }>(API_ENDPOINTS.anomalies.resolve(id), {
      method: 'PUT',
      body: JSON.stringify({ resolution_note: resolutionNote }),
    }),
};

export const statisticsApi = {
  dashboard: () => fetchApi<StatisticsDashboard>(API_ENDPOINTS.statistics.dashboard),
  completeness: (params?: Record<string, string>) => {
    const query = params ? `?${new URLSearchParams(params)}` : '';
    return fetchApi<ApiResponse<unknown>>(`${API_ENDPOINTS.statistics.completeness}${query}`);
  },
  rectificationRate: (params?: Record<string, string>) => {
    const query = params ? `?${new URLSearchParams(params)}` : '';
    return fetchApi<ApiResponse<unknown>>(`${API_ENDPOINTS.statistics.rectificationRate}${query}`);
  },
  retestPassRate: (params?: Record<string, string>) => {
    const query = params ? `?${new URLSearchParams(params)}` : '';
    return fetchApi<ApiResponse<unknown>>(`${API_ENDPOINTS.statistics.retestPassRate}${query}`);
  },
};

export const calculateApi = {
  priority: (data: {
    complaint_level: string;
    sampling_time_compliant: boolean;
    reading_deviation: number;
    evidence_completeness: number;
  }) =>
    fetchApi<{ priority: number; level: string; explanation: string }>(
      API_ENDPOINTS.calculate.priority,
      {
        method: 'POST',
        body: JSON.stringify(data),
      }
    ),
  completeness: (data: {
    complaint_id: number;
    sampling_point_count: number;
    reading_count: number;
    evidence_count: number;
  }) =>
    fetchApi<{
      completeness: number;
      required_fields: string[];
      missing_fields: string[];
      is_complete: boolean;
      explanation: string;
    }>(API_ENDPOINTS.calculate.completeness, {
      method: 'POST',
      body: JSON.stringify(data),
    }),
  compliance: (data: { measurement_time: string; period_type?: string }) =>
    fetchApi<{
      is_compliant: boolean;
      period_type: string;
      time_from: string;
      time_to: string;
      violation_msg: string;
    }>(API_ENDPOINTS.calculate.compliance, {
      method: 'POST',
      body: JSON.stringify(data),
    }),
};

export const healthApi = {
  check: () => fetchApi<{ status: string; message: string }>(API_ENDPOINTS.health),
};

export const seedApi = {
  reset: () =>
    fetchApi<{ message: string }>(API_ENDPOINTS.seed.reset, { method: 'POST' }),
  browse: () => fetchApi<Record<string, unknown>>(API_ENDPOINTS.seed.browse),
};

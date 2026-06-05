export const API_BASE_URL = '/api';

export const API_ENDPOINTS = {
  complaints: {
    list: `${API_BASE_URL}/complaints`,
    create: `${API_BASE_URL}/complaints`,
    get: (id: number) => `${API_BASE_URL}/complaints/${id}`,
    update: (id: number) => `${API_BASE_URL}/complaints/${id}`,
    delete: (id: number) => `${API_BASE_URL}/complaints/${id}`,
    transition: (id: number) => `${API_BASE_URL}/complaints/${id}/transition`,
  },
  samplingPoints: {
    list: `${API_BASE_URL}/sampling-points`,
    create: `${API_BASE_URL}/sampling-points`,
    get: (id: number) => `${API_BASE_URL}/sampling-points/${id}`,
    update: (id: number) => `${API_BASE_URL}/sampling-points/${id}`,
    delete: (id: number) => `${API_BASE_URL}/sampling-points/${id}`,
    transition: (id: number) => `${API_BASE_URL}/sampling-points/${id}/transition`,
    validate: (id: number) => `${API_BASE_URL}/sampling-points/${id}/validate`,
  },
  noiseReadings: {
    list: `${API_BASE_URL}/noise-readings`,
    create: `${API_BASE_URL}/noise-readings`,
    get: (id: number) => `${API_BASE_URL}/noise-readings/${id}`,
    update: (id: number) => `${API_BASE_URL}/noise-readings/${id}`,
    delete: (id: number) => `${API_BASE_URL}/noise-readings/${id}`,
    byPoint: (id: number) => `${API_BASE_URL}/noise-readings/by-point/${id}`,
  },
  timePeriods: {
    list: `${API_BASE_URL}/time-periods`,
    create: `${API_BASE_URL}/time-periods`,
    get: (id: number) => `${API_BASE_URL}/time-periods/${id}`,
    update: (id: number) => `${API_BASE_URL}/time-periods/${id}`,
    delete: (id: number) => `${API_BASE_URL}/time-periods/${id}`,
    batchImport: `${API_BASE_URL}/time-periods/batch-import`,
  },
  anomalies: {
    list: `${API_BASE_URL}/anomalies`,
    get: (id: number) => `${API_BASE_URL}/anomalies/${id}`,
    resolve: (id: number) => `${API_BASE_URL}/anomalies/${id}/resolve`,
  },
  calculate: {
    priority: `${API_BASE_URL}/calculate/priority`,
    completeness: `${API_BASE_URL}/calculate/completeness`,
    compliance: `${API_BASE_URL}/calculate/compliance`,
  },
  statistics: {
    dashboard: `${API_BASE_URL}/statistics/dashboard`,
    completeness: `${API_BASE_URL}/statistics/completeness`,
    rectificationRate: `${API_BASE_URL}/statistics/rectification-rate`,
    retestPassRate: `${API_BASE_URL}/statistics/retest-pass-rate`,
  },
  archive: {
    snapshot: `${API_BASE_URL}/archive/snapshot`,
    snapshots: `${API_BASE_URL}/archive/snapshots`,
    export: `${API_BASE_URL}/archive/export`,
  },
  auditLogs: {
    list: `${API_BASE_URL}/audit-logs`,
  },
  rules: {
    list: `${API_BASE_URL}/rules`,
    create: `${API_BASE_URL}/rules`,
    get: (id: number) => `${API_BASE_URL}/rules/${id}`,
    update: (id: number) => `${API_BASE_URL}/rules/${id}`,
  },
  health: `${API_BASE_URL}/health`,
  seed: {
    reset: `${API_BASE_URL}/seed/reset`,
    browse: `${API_BASE_URL}/seed/browse`,
  },
} as const;

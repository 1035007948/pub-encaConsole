import { create } from 'zustand';
import { AnomalyEvent, anomaliesApi } from '../api';

interface AnomalyState {
  anomalies: AnomalyEvent[];
  currentAnomaly: AnomalyEvent | null;
  total: number;
  loading: boolean;
  error: string | null;
  fetchAnomalies: (params?: Record<string, string>) => Promise<void>;
  fetchAnomaly: (id: number) => Promise<void>;
  resolveAnomaly: (id: number, resolutionNote: string) => Promise<void>;
  setCurrentAnomaly: (anomaly: AnomalyEvent | null) => void;
  clearError: () => void;
}

export const useAnomalyStore = create<AnomalyState>((set, get) => ({
  anomalies: [],
  currentAnomaly: null,
  total: 0,
  loading: false,
  error: null,

  fetchAnomalies: async (params) => {
    set({ loading: true, error: null });
    try {
      const response = await anomaliesApi.list(params);
      set({
        anomalies: response.items || [],
        total: response.total || 0,
        loading: false,
      });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  fetchAnomaly: async (id) => {
    set({ loading: true, error: null });
    try {
      const anomaly = await anomaliesApi.get(id);
      set({ currentAnomaly: anomaly, loading: false });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  resolveAnomaly: async (id, resolutionNote) => {
    set({ loading: true, error: null });
    try {
      await anomaliesApi.resolve(id, resolutionNote);
      await get().fetchAnomalies();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  setCurrentAnomaly: (anomaly) => set({ currentAnomaly: anomaly }),
  clearError: () => set({ error: null }),
}));

import { create } from 'zustand';
import { SamplingPoint, samplingPointsApi } from '../api';

interface SamplingPointState {
  samplingPoints: SamplingPoint[];
  currentPoint: SamplingPoint | null;
  total: number;
  loading: boolean;
  error: string | null;
  fetchSamplingPoints: (params?: Record<string, string>) => Promise<void>;
  fetchSamplingPoint: (id: number) => Promise<void>;
  createSamplingPoint: (data: Partial<SamplingPoint>) => Promise<void>;
  updateSamplingPoint: (id: number, data: Partial<SamplingPoint>) => Promise<void>;
  deleteSamplingPoint: (id: number) => Promise<void>;
  transitionSamplingPoint: (id: number, toStatus: string, reason?: string) => Promise<void>;
  validateConsistency: (id: number) => Promise<{ isValid: boolean; missingFields: string[] }>;
  setCurrentPoint: (point: SamplingPoint | null) => void;
  clearError: () => void;
}

export const useSamplingPointStore = create<SamplingPointState>((set, get) => ({
  samplingPoints: [],
  currentPoint: null,
  total: 0,
  loading: false,
  error: null,

  fetchSamplingPoints: async (params) => {
    set({ loading: true, error: null });
    try {
      const response = await samplingPointsApi.list(params);
      set({
        samplingPoints: response.items || [],
        total: response.total || 0,
        loading: false,
      });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  fetchSamplingPoint: async (id) => {
    set({ loading: true, error: null });
    try {
      const point = await samplingPointsApi.get(id);
      set({ currentPoint: point, loading: false });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  createSamplingPoint: async (data) => {
    set({ loading: true, error: null });
    try {
      await samplingPointsApi.create(data);
      await get().fetchSamplingPoints();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  updateSamplingPoint: async (id, data) => {
    set({ loading: true, error: null });
    try {
      await samplingPointsApi.update(id, data);
      await get().fetchSamplingPoints();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  deleteSamplingPoint: async (id) => {
    set({ loading: true, error: null });
    try {
      await samplingPointsApi.delete(id);
      await get().fetchSamplingPoints();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  transitionSamplingPoint: async (id, toStatus, reason) => {
    set({ loading: true, error: null });
    try {
      await samplingPointsApi.transition(id, toStatus, reason);
      await get().fetchSamplingPoints();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  validateConsistency: async (id) => {
    try {
      const result = await samplingPointsApi.validate(id);
      return { isValid: result.is_valid, missingFields: result.missing_fields };
    } catch (error) {
      return { isValid: false, missingFields: [] };
    }
  },

  setCurrentPoint: (point) => set({ currentPoint: point }),
  clearError: () => set({ error: null }),
}));

import { create } from 'zustand';
import { NoiseReading, noiseReadingsApi } from '../api';

interface NoiseReadingState {
  readings: NoiseReading[];
  currentReading: NoiseReading | null;
  total: number;
  loading: boolean;
  error: string | null;
  fetchReadings: (params?: Record<string, string>) => Promise<void>;
  fetchReading: (id: number) => Promise<void>;
  fetchByPoint: (pointId: number) => Promise<void>;
  createReading: (data: Partial<NoiseReading>) => Promise<void>;
  updateReading: (id: number, data: Partial<NoiseReading>) => Promise<void>;
  deleteReading: (id: number) => Promise<void>;
  setCurrentReading: (reading: NoiseReading | null) => void;
  clearError: () => void;
}

export const useNoiseReadingStore = create<NoiseReadingState>((set, get) => ({
  readings: [],
  currentReading: null,
  total: 0,
  loading: false,
  error: null,

  fetchReadings: async (params) => {
    set({ loading: true, error: null });
    try {
      const response = await noiseReadingsApi.list(params);
      set({
        readings: response.items || [],
        total: response.total || 0,
        loading: false,
      });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  fetchReading: async (id) => {
    set({ loading: true, error: null });
    try {
      const reading = await noiseReadingsApi.get(id);
      set({ currentReading: reading, loading: false });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  fetchByPoint: async (pointId) => {
    set({ loading: true, error: null });
    try {
      const response = await noiseReadingsApi.byPoint(pointId);
      set({
        readings: response.items || [],
        total: response.total || 0,
        loading: false,
      });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  createReading: async (data) => {
    set({ loading: true, error: null });
    try {
      await noiseReadingsApi.create(data);
      await get().fetchReadings();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  updateReading: async (id, data) => {
    set({ loading: true, error: null });
    try {
      await noiseReadingsApi.update(id, data);
      await get().fetchReadings();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  deleteReading: async (id) => {
    set({ loading: true, error: null });
    try {
      await noiseReadingsApi.delete(id);
      await get().fetchReadings();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  setCurrentReading: (reading) => set({ currentReading: reading }),
  clearError: () => set({ error: null }),
}));

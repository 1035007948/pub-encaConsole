import { create } from 'zustand';
import { StatisticsDashboard, statisticsApi } from '../api';

interface StatisticsState {
  dashboard: StatisticsDashboard | null;
  loading: boolean;
  error: string | null;
  fetchDashboard: () => Promise<void>;
  clearError: () => void;
}

export const useStatisticsStore = create<StatisticsState>((set) => ({
  dashboard: null,
  loading: false,
  error: null,

  fetchDashboard: async () => {
    set({ loading: true, error: null });
    try {
      const dashboard = await statisticsApi.dashboard();
      set({ dashboard, loading: false });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  clearError: () => set({ error: null }),
}));

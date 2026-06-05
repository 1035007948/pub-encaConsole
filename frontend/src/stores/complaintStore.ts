import { create } from 'zustand';
import { Complaint, complaintsApi } from '../api';

interface ComplaintState {
  complaints: Complaint[];
  currentComplaint: Complaint | null;
  total: number;
  loading: boolean;
  error: string | null;
  fetchComplaints: (params?: Record<string, string>) => Promise<void>;
  fetchComplaint: (id: number) => Promise<void>;
  createComplaint: (data: Partial<Complaint>) => Promise<void>;
  updateComplaint: (id: number, data: Partial<Complaint>) => Promise<void>;
  deleteComplaint: (id: number) => Promise<void>;
  transitionComplaint: (id: number, toStatus: string, reason?: string) => Promise<void>;
  setCurrentComplaint: (complaint: Complaint | null) => void;
  clearError: () => void;
}

export const useComplaintStore = create<ComplaintState>((set, get) => ({
  complaints: [],
  currentComplaint: null,
  total: 0,
  loading: false,
  error: null,

  fetchComplaints: async (params) => {
    set({ loading: true, error: null });
    try {
      const response = await complaintsApi.list(params);
      set({
        complaints: response.items || [],
        total: response.total || 0,
        loading: false,
      });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  fetchComplaint: async (id) => {
    set({ loading: true, error: null });
    try {
      const complaint = await complaintsApi.get(id);
      set({ currentComplaint: complaint, loading: false });
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  createComplaint: async (data) => {
    set({ loading: true, error: null });
    try {
      await complaintsApi.create(data);
      await get().fetchComplaints();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  updateComplaint: async (id, data) => {
    set({ loading: true, error: null });
    try {
      await complaintsApi.update(id, data);
      await get().fetchComplaints();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  deleteComplaint: async (id) => {
    set({ loading: true, error: null });
    try {
      await complaintsApi.delete(id);
      await get().fetchComplaints();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  transitionComplaint: async (id, toStatus, reason) => {
    set({ loading: true, error: null });
    try {
      await complaintsApi.transition(id, toStatus, reason);
      await get().fetchComplaints();
    } catch (error) {
      set({ error: (error as Error).message, loading: false });
    }
  },

  setCurrentComplaint: (complaint) => set({ currentComplaint: complaint }),
  clearError: () => set({ error: null }),
}));

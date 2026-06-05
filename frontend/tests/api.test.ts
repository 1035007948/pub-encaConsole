import { describe, it, expect } from 'vitest';
import { API_ENDPOINTS } from '../src/api/config';

describe('API Endpoints', () => {
  it('should have correct complaint endpoints', () => {
    expect(API_ENDPOINTS.complaints.list).toBe('/api/complaints');
    expect(API_ENDPOINTS.complaints.create).toBe('/api/complaints');
    expect(API_ENDPOINTS.complaints.get(1)).toBe('/api/complaints/1');
    expect(API_ENDPOINTS.complaints.update(1)).toBe('/api/complaints/1');
    expect(API_ENDPOINTS.complaints.delete(1)).toBe('/api/complaints/1');
    expect(API_ENDPOINTS.complaints.transition(1)).toBe('/api/complaints/1/transition');
  });

  it('should have correct sampling point endpoints', () => {
    expect(API_ENDPOINTS.samplingPoints.list).toBe('/api/sampling-points');
    expect(API_ENDPOINTS.samplingPoints.validate(1)).toBe('/api/sampling-points/1/validate');
  });

  it('should have correct noise reading endpoints', () => {
    expect(API_ENDPOINTS.noiseReadings.list).toBe('/api/noise-readings');
    expect(API_ENDPOINTS.noiseReadings.byPoint(1)).toBe('/api/noise-readings/by-point/1');
  });

  it('should have correct time period endpoints', () => {
    expect(API_ENDPOINTS.timePeriods.list).toBe('/api/time-periods');
    expect(API_ENDPOINTS.timePeriods.batchImport).toBe('/api/time-periods/batch-import');
  });

  it('should have correct anomaly endpoints', () => {
    expect(API_ENDPOINTS.anomalies.list).toBe('/api/anomalies');
    expect(API_ENDPOINTS.anomalies.resolve(1)).toBe('/api/anomalies/1/resolve');
  });

  it('should have correct calculate endpoints', () => {
    expect(API_ENDPOINTS.calculate.priority).toBe('/api/calculate/priority');
    expect(API_ENDPOINTS.calculate.completeness).toBe('/api/calculate/completeness');
    expect(API_ENDPOINTS.calculate.compliance).toBe('/api/calculate/compliance');
  });

  it('should have correct statistics endpoints', () => {
    expect(API_ENDPOINTS.statistics.dashboard).toBe('/api/statistics/dashboard');
    expect(API_ENDPOINTS.statistics.completeness).toBe('/api/statistics/completeness');
  });

  it('should have correct archive endpoints', () => {
    expect(API_ENDPOINTS.archive.snapshot).toBe('/api/archive/snapshot');
    expect(API_ENDPOINTS.archive.export).toBe('/api/archive/export');
  });

  it('should have health and seed endpoints', () => {
    expect(API_ENDPOINTS.health).toBe('/api/health');
    expect(API_ENDPOINTS.seed.reset).toBe('/api/seed/reset');
    expect(API_ENDPOINTS.seed.browse).toBe('/api/seed/browse');
  });
});

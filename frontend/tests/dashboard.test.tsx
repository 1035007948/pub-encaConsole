import { describe, it, expect, vi } from 'vitest';
import { render, screen } from '@testing-library/react';
import { BrowserRouter } from 'react-router-dom';
import { MantineProvider, createTheme } from '@mantine/core';
import { StatisticsDashboard } from '../src/pages/StatisticsDashboard';

const theme = createTheme();

const renderWithProviders = (component: React.ReactNode) => {
  return render(
    <MantineProvider theme={theme}>
      <BrowserRouter>{component}</BrowserRouter>
    </MantineProvider>
  );
};

describe('StatisticsDashboard', () => {
  it('renders dashboard title', () => {
    renderWithProviders(<StatisticsDashboard />);
    expect(screen.getByText('证据完整度统计驾驶舱')).toBeInTheDocument();
  });
});

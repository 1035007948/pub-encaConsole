import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { MantineProvider, createTheme } from '@mantine/core';
import { Notifications } from '@mantine/notifications';
import { AppLayout } from './components/AppLayout';
import { StatisticsDashboard } from './pages/StatisticsDashboard';
import { ComplaintListPage } from './pages/ComplaintListPage';
import { SamplingPointWorkbench } from './pages/SamplingPointWorkbench';
import { NoiseReadingPage } from './pages/NoiseReadingPage';
import { AnomalyTriagePage } from './pages/AnomalyTriagePage';
import '@mantine/core/styles.css';
import '@mantine/dates/styles.css';
import '@mantine/notifications/styles.css';

const theme = createTheme({
  primaryColor: 'blue',
  fontFamily: 'Inter, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, sans-serif',
  headings: {
    fontFamily: 'Inter, -apple-system, BlinkMacSystemFont, Segoe UI, Roboto, sans-serif',
  },
});

export function App() {
  return (
    <MantineProvider theme={theme}>
      <Notifications position="top-right" />
      <BrowserRouter>
        <AppLayout>
          <Routes>
            <Route path="/" element={<StatisticsDashboard />} />
            <Route path="/complaints" element={<ComplaintListPage />} />
            <Route path="/sampling-points" element={<SamplingPointWorkbench />} />
            <Route path="/noise-readings" element={<NoiseReadingPage />} />
            <Route path="/anomalies" element={<AnomalyTriagePage />} />
            <Route path="/rules" element={<RulesPage />} />
            <Route path="/seed" element={<SeedDataPage />} />
          </Routes>
        </AppLayout>
      </BrowserRouter>
    </MantineProvider>
  );
}

function RulesPage() {
  return (
    <div>
      <h1>规则配置</h1>
      <p>规则配置页面开发中...</p>
    </div>
  );
}

function SeedDataPage() {
  return (
    <div>
      <h1>Seed数据浏览</h1>
      <p>Seed数据浏览页面开发中...</p>
    </div>
  );
}

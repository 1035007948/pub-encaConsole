import { useEffect } from 'react';
import { Stack, Paper, Group, Text, Grid, Progress, Badge } from '@mantine/core';
import { useStatisticsStore } from '../stores/statisticsStore';
import { ChartWidget } from '../components/ChartWidget';

export function StatisticsDashboard() {
  const { dashboard, loading, fetchDashboard } = useStatisticsStore();

  useEffect(() => {
    fetchDashboard();
  }, [fetchDashboard]);

  if (loading || !dashboard) {
    return (
      <Stack gap="md">
        <Text size="xl" fw={700}>
          统计驾驶舱
        </Text>
        <Text c="dimmed">加载中...</Text>
      </Stack>
    );
  }

  return (
    <Stack gap="md">
      <Text size="xl" fw={700}>
        证据完整度统计驾驶舱
      </Text>

      <Grid>
        <Grid.Col span={3}>
          <ChartWidget
            title="投诉单总数"
            type="metric"
            value={dashboard.total_complaints}
            unit="件"
            height={150}
          />
        </Grid.Col>
        <Grid.Col span={3}>
          <ChartWidget
            title="待处理投诉"
            type="metric"
            value={dashboard.pending_complaints}
            unit="件"
            height={150}
          />
        </Grid.Col>
        <Grid.Col span={3}>
          <ChartWidget
            title="已完成投诉"
            type="metric"
            value={dashboard.completed_complaints}
            unit="件"
            height={150}
          />
        </Grid.Col>
        <Grid.Col span={3}>
          <ChartWidget
            title="异常事件"
            type="metric"
            value={dashboard.open_anomalies}
            unit="件"
            height={150}
          />
        </Grid.Col>
      </Grid>

      <Grid>
        <Grid.Col span={6}>
          <Paper shadow="sm" p="md" withBorder>
            <Stack gap="md">
              <Text fw={500}>采样统计</Text>
              <Group justify="space-between">
                <Text size="sm">采样点位</Text>
                <Text size="sm" fw={500}>
                  {dashboard.total_sampling_points} 个
                </Text>
              </Group>
              <Group justify="space-between">
                <Text size="sm">噪声读数</Text>
                <Text size="sm" fw={500}>
                  {dashboard.total_noise_readings} 条
                </Text>
              </Group>
              <Group justify="space-between">
                <Text size="sm">超标读数</Text>
                <Text size="sm" fw={500} c="red">
                  {dashboard.exceeded_readings} 条
                </Text>
              </Group>
              <Group justify="space-between">
                <Text size="sm">平均等效声级</Text>
                <Text size="sm" fw={500}>
                  {dashboard.average_leq.toFixed(1)} dB
                </Text>
              </Group>
            </Stack>
          </Paper>
        </Grid.Col>

        <Grid.Col span={6}>
          <Paper shadow="sm" p="md" withBorder>
            <Stack gap="md">
              <Text fw={500}>质量指标</Text>
              <Stack gap="xs">
                <Group justify="space-between">
                  <Text size="sm">证据完整度</Text>
                  <Badge>
                    {(dashboard.evidence_completeness * 100).toFixed(1)}%
                  </Badge>
                </Group>
                <Progress
                  value={dashboard.evidence_completeness * 100}
                  color="blue"
                  size="sm"
                />
              </Stack>
              <Stack gap="xs">
                <Group justify="space-between">
                  <Text size="sm">整改闭环率</Text>
                  <Badge>
                    {(dashboard.rectification_rate * 100).toFixed(1)}%
                  </Badge>
                </Group>
                <Progress
                  value={dashboard.rectification_rate * 100}
                  color="green"
                  size="sm"
                />
              </Stack>
              <Stack gap="xs">
                <Group justify="space-between">
                  <Text size="sm">复测通过率</Text>
                  <Badge>
                    {(dashboard.retest_pass_rate * 100).toFixed(1)}%
                  </Badge>
                </Group>
                <Progress
                  value={dashboard.retest_pass_rate * 100}
                  color="orange"
                  size="sm"
                />
              </Stack>
            </Stack>
          </Paper>
        </Grid.Col>
      </Grid>

      <Grid>
        <Grid.Col span={6}>
          <ChartWidget
            title="投诉状态分布"
            type="bar"
            data={[
              { label: '待处理', value: dashboard.pending_complaints },
              { label: '已完成', value: dashboard.completed_complaints },
              { label: '异常', value: dashboard.open_anomalies },
            ]}
            height={200}
          />
        </Grid.Col>
        <Grid.Col span={6}>
          <ChartWidget
            title="读数达标情况"
            type="pie"
            data={[
              {
                label: '达标',
                value: dashboard.total_noise_readings - dashboard.exceeded_readings,
                color: '#40c057',
              },
              { label: '超标', value: dashboard.exceeded_readings, color: '#fa5252' },
            ]}
            height={200}
          />
        </Grid.Col>
      </Grid>
    </Stack>
  );
}

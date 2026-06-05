import { useEffect, useState } from 'react';
import {
  Stack,
  Paper,
  Group,
  Button,
  Text,
  Badge,
  Grid,
  Tabs,
  Code,
  JsonInput,
} from '@mantine/core';
import { IconRefresh, IconDatabase, IconCheck } from '@tabler/icons-react';
import { seedApi, healthApi } from '../api';

export function SeedDataPage() {
  const [seedData, setSeedData] = useState<Record<string, unknown> | null>(null);
  const [healthStatus, setHealthStatus] = useState<{ status: string; message: string } | null>(
    null
  );
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    fetchSeedData();
    checkHealth();
  }, []);

  const fetchSeedData = async () => {
    setLoading(true);
    try {
      const data = await seedApi.browse();
      setSeedData(data);
    } catch (error) {
      console.error('Failed to fetch seed data:', error);
    } finally {
      setLoading(false);
    }
  };

  const checkHealth = async () => {
    try {
      const status = await healthApi.check();
      setHealthStatus(status);
    } catch (error) {
      setHealthStatus({ status: 'error', message: '无法连接到后端服务' });
    }
  };

  const handleReset = async () => {
    try {
      await seedApi.reset();
      fetchSeedData();
    } catch (error) {
      console.error('Failed to reset seed data:', error);
    }
  };

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Text size="xl" fw={700}>
          Seed数据浏览与接口健康检查
        </Text>
        <Group>
          <Button
            variant="light"
            leftSection={<IconRefresh size={16} />}
            onClick={() => {
              fetchSeedData();
              checkHealth();
            }}
          >
            刷新
          </Button>
          <Button
            color="orange"
            leftSection={<IconDatabase size={16} />}
            onClick={handleReset}
          >
            重置Seed数据
          </Button>
        </Group>
      </Group>

      <Grid>
        <Grid.Col span={12}>
          <Paper shadow="sm" p="md" withBorder>
            <Stack gap="sm">
              <Group justify="space-between">
                <Text fw={500}>接口健康状态</Text>
                {healthStatus && (
                  <Badge
                    size="lg"
                    color={healthStatus.status === 'ok' ? 'green' : 'red'}
                    leftSection={
                      healthStatus.status === 'ok' ? (
                        <IconCheck size={12} />
                      ) : null
                    }
                  >
                    {healthStatus.status === 'ok' ? '正常' : '异常'}
                  </Badge>
                )}
              </Group>
              {healthStatus && (
                <Text size="sm" c="dimmed">
                  {healthStatus.message}
                </Text>
              )}
            </Stack>
          </Paper>
        </Grid.Col>
      </Grid>

      <Paper shadow="sm" withBorder>
        <Tabs defaultValue="overview">
          <Tabs.List>
            <Tabs.Tab value="overview">概览</Tabs.Tab>
            <Tabs.Tab value="complaints">投诉单</Tabs.Tab>
            <Tabs.Tab value="sampling-points">采样点位</Tabs.Tab>
            <Tabs.Tab value="readings">噪声读数</Tabs.Tab>
            <Tabs.Tab value="anomalies">异常事件</Tabs.Tab>
          </Tabs.List>

          <Tabs.Panel value="overview" p="md">
            <Stack gap="md">
              <Text fw={500}>数据统计</Text>
              {seedData && (
                <Grid>
                  <Grid.Col span={3}>
                    <Paper shadow="xs" p="sm" withBorder>
                      <Text size="sm" c="dimmed">
                        投诉单
                      </Text>
                      <Text size="xl" fw={700}>
                        {Array.isArray(seedData['complaints'])
                          ? seedData['complaints'].length
                          : 0}
                      </Text>
                    </Paper>
                  </Grid.Col>
                  <Grid.Col span={3}>
                    <Paper shadow="xs" p="sm" withBorder>
                      <Text size="sm" c="dimmed">
                        采样点位
                      </Text>
                      <Text size="xl" fw={700}>
                        {Array.isArray(seedData['sampling_points'])
                          ? seedData['sampling_points'].length
                          : 0}
                      </Text>
                    </Paper>
                  </Grid.Col>
                  <Grid.Col span={3}>
                    <Paper shadow="xs" p="sm" withBorder>
                      <Text size="sm" c="dimmed">
                        噪声读数
                      </Text>
                      <Text size="xl" fw={700}>
                        {Array.isArray(seedData['noise_readings'])
                          ? seedData['noise_readings'].length
                          : 0}
                      </Text>
                    </Paper>
                  </Grid.Col>
                  <Grid.Col span={3}>
                    <Paper shadow="xs" p="sm" withBorder>
                      <Text size="sm" c="dimmed">
                        异常事件
                      </Text>
                      <Text size="xl" fw={700}>
                        {Array.isArray(seedData['anomalies'])
                          ? seedData['anomalies'].length
                          : 0}
                      </Text>
                    </Paper>
                  </Grid.Col>
                </Grid>
              )}
            </Stack>
          </Tabs.Panel>

          <Tabs.Panel value="complaints" p="md">
            <Stack gap="md">
              <Text fw={500}>投诉单数据</Text>
              {seedData && seedData['complaints'] && (
                <JsonInput
                  value={JSON.stringify(seedData['complaints'], null, 2)}
                  rows={15}
                  readOnly
                />
              )}
            </Stack>
          </Tabs.Panel>

          <Tabs.Panel value="sampling-points" p="md">
            <Stack gap="md">
              <Text fw={500}>采样点位数据</Text>
              {seedData && seedData['sampling_points'] && (
                <JsonInput
                  value={JSON.stringify(seedData['sampling_points'], null, 2)}
                  rows={15}
                  readOnly
                />
              )}
            </Stack>
          </Tabs.Panel>

          <Tabs.Panel value="readings" p="md">
            <Stack gap="md">
              <Text fw={500}>噪声读数数据</Text>
              {seedData && seedData['noise_readings'] && (
                <JsonInput
                  value={JSON.stringify(seedData['noise_readings'], null, 2)}
                  rows={15}
                  readOnly
                />
              )}
            </Stack>
          </Tabs.Panel>

          <Tabs.Panel value="anomalies" p="md">
            <Stack gap="md">
              <Text fw={500}>异常事件数据</Text>
              {seedData && seedData['anomalies'] && (
                <JsonInput
                  value={JSON.stringify(seedData['anomalies'], null, 2)}
                  rows={15}
                  readOnly
                />
              )}
            </Stack>
          </Tabs.Panel>
        </Tabs>
      </Paper>
    </Stack>
  );
}

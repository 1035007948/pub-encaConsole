import { useEffect, useState } from 'react';
import { Stack, Paper, Group, Button, Text, Badge, Grid, Modal, TextInput, NumberInput } from '@mantine/core';
import { IconPlus, IconMapPin, IconCalendar } from '@tabler/icons-react';
import { useSamplingPointStore } from '../stores/samplingPointStore';
import { ChartWidget } from '../components/ChartWidget';
import { SamplingPoint } from '../api';

export function SamplingPointWorkbench() {
  const {
    samplingPoints,
    loading,
    fetchSamplingPoints,
    createSamplingPoint,
    currentPoint,
    setCurrentPoint,
  } = useSamplingPointStore();

  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [formData, setFormData] = useState<Partial<SamplingPoint>>({});

  useEffect(() => {
    fetchSamplingPoints();
  }, [fetchSamplingPoints]);

  const statusCounts = samplingPoints.reduce(
    (acc, point) => {
      acc[point.status] = (acc[point.status] || 0) + 1;
      return acc;
    },
    {} as Record<string, number>
  );

  const handleCreate = async () => {
    await createSamplingPoint(formData);
    setCreateModalOpen(false);
    setFormData({});
  };

  const getStatusColor = (status: string) => {
    const colorMap: Record<string, string> = {
      draft: 'gray',
      pending: 'yellow',
      scheduled: 'blue',
      in_progress: 'cyan',
      completed: 'green',
      rejected: 'red',
    };
    return colorMap[status] || 'gray';
  };

  const getStatusLabel = (status: string) => {
    const labelMap: Record<string, string> = {
      draft: '草稿',
      pending: '待安排',
      scheduled: '已安排',
      in_progress: '进行中',
      completed: '已完成',
      rejected: '已驳回',
    };
    return labelMap[status] || status;
  };

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Text size="xl" fw={700}>
          采样点位工作台
        </Text>
        <Button
          leftSection={<IconPlus size={16} />}
          onClick={() => setCreateModalOpen(true)}
        >
          安排采样点位
        </Button>
      </Group>

      <Grid>
        <Grid.Col span={4}>
          <ChartWidget
            title="点位状态分布"
            type="pie"
            data={Object.entries(statusCounts).map(([label, value]) => ({
              label: getStatusLabel(label),
              value,
            }))}
            height={200}
          />
        </Grid.Col>
        <Grid.Col span={4}>
          <ChartWidget
            title="总点位数"
            type="metric"
            value={samplingPoints.length}
            unit="个"
            height={200}
          />
        </Grid.Col>
        <Grid.Col span={4}>
          <ChartWidget
            title="待安排点位"
            type="metric"
            value={statusCounts['pending'] || 0}
            unit="个"
            height={200}
          />
        </Grid.Col>
      </Grid>

      <Paper shadow="sm" withBorder p="md">
        <Stack gap="md">
          <Text fw={500}>点位列表</Text>
          <Grid>
            {samplingPoints.map((point) => (
              <Grid.Col key={point.id} span={4}>
                <Paper
                  shadow="xs"
                  p="sm"
                  withBorder
                  style={{ cursor: 'pointer' }}
                  onClick={() => setCurrentPoint(point)}
                >
                  <Stack gap="xs">
                    <Group justify="space-between">
                      <Text size="sm" fw={500}>
                        {point.point_no}
                      </Text>
                      <Badge size="sm" color={getStatusColor(point.status)}>
                        {getStatusLabel(point.status)}
                      </Badge>
                    </Group>
                    <Text size="xs" c="dimmed">
                      {point.point_name}
                    </Text>
                    <Group gap="xs">
                      <IconMapPin size={12} />
                      <Text size="xs" c="dimmed">
                        {point.address}
                      </Text>
                    </Group>
                    {point.scheduled_date && (
                      <Group gap="xs">
                        <IconCalendar size={12} />
                        <Text size="xs" c="dimmed">
                          {point.scheduled_date} {point.scheduled_time_from}-{point.scheduled_time_to}
                        </Text>
                      </Group>
                    )}
                  </Stack>
                </Paper>
              </Grid.Col>
            ))}
          </Grid>
        </Stack>
      </Paper>

      <Modal
        opened={createModalOpen}
        onClose={() => setCreateModalOpen(false)}
        title="安排采样点位"
        size="lg"
      >
        <Stack gap="md">
          <TextInput
            label="点位编号"
            required
            value={formData.point_no || ''}
            onChange={(e) => setFormData({ ...formData, point_no: e.target.value })}
          />
          <TextInput
            label="点位名称"
            required
            value={formData.point_name || ''}
            onChange={(e) => setFormData({ ...formData, point_name: e.target.value })}
          />
          <TextInput
            label="地址"
            required
            value={formData.address || ''}
            onChange={(e) => setFormData({ ...formData, address: e.target.value })}
          />
          <Grid>
            <Grid.Col span={6}>
              <NumberInput
                label="经度"
                decimalScale={6}
                value={formData.longitude}
                onChange={(value) => setFormData({ ...formData, longitude: Number(value) })}
              />
            </Grid.Col>
            <Grid.Col span={6}>
              <NumberInput
                label="纬度"
                decimalScale={6}
                value={formData.latitude}
                onChange={(value) => setFormData({ ...formData, latitude: Number(value) })}
              />
            </Grid.Col>
          </Grid>
          <TextInput
            label="负责人"
            value={formData.responsible_user || ''}
            onChange={(e) => setFormData({ ...formData, responsible_user: e.target.value })}
          />
          <Group justify="flex-end">
            <Button variant="subtle" onClick={() => setCreateModalOpen(false)}>
              取消
            </Button>
            <Button onClick={handleCreate}>创建</Button>
          </Group>
        </Stack>
      </Modal>

      <Modal
        opened={currentPoint !== null}
        onClose={() => setCurrentPoint(null)}
        title="点位详情"
        size="lg"
      >
        {currentPoint && (
          <Stack gap="md">
            <Group justify="space-between">
              <Text fw={500}>{currentPoint.point_no}</Text>
              <Badge color={getStatusColor(currentPoint.status)}>
                {getStatusLabel(currentPoint.status)}
              </Badge>
            </Group>
            <Text size="sm" c="dimmed">
              {currentPoint.point_name}
            </Text>
            <Text size="sm">{currentPoint.address}</Text>
            <Text size="sm" c="dimmed">
              坐标: {currentPoint.longitude}, {currentPoint.latitude}
            </Text>
            {currentPoint.scheduled_date && (
              <Text size="sm">
                安排时间: {currentPoint.scheduled_date} {currentPoint.scheduled_time_from}-
                {currentPoint.scheduled_time_to}
              </Text>
            )}
          </Stack>
        )}
      </Modal>
    </Stack>
  );
}

import { useEffect, useState } from 'react';
import {
  Stack,
  Paper,
  Group,
  Button,
  Text,
  Badge,
  Modal,
  TextInput,
  Select,
  NumberInput,
  Textarea,
} from '@mantine/core';
import { IconPlus, IconRefresh } from '@tabler/icons-react';
import { timePeriodsApi, TimePeriod } from '../api';

export function TimePeriodPage() {
  const [timePeriods, setTimePeriods] = useState<TimePeriod[]>([]);
  const [loading, setLoading] = useState(false);
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [formData, setFormData] = useState<Partial<TimePeriod>>({});

  useEffect(() => {
    fetchTimePeriods();
  }, []);

  const fetchTimePeriods = async () => {
    setLoading(true);
    try {
      const response = await timePeriodsApi.list();
      setTimePeriods(response.items || []);
    } catch (error) {
      console.error('Failed to fetch time periods:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleCreate = async () => {
    try {
      await timePeriodsApi.create(formData);
      setCreateModalOpen(false);
      setFormData({});
      fetchTimePeriods();
    } catch (error) {
      console.error('Failed to create time period:', error);
    }
  };

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Text size="xl" fw={700}>
          时段分类管理
        </Text>
        <Group>
          <Button
            variant="light"
            leftSection={<IconRefresh size={16} />}
            onClick={fetchTimePeriods}
          >
            刷新
          </Button>
          <Button
            leftSection={<IconPlus size={16} />}
            onClick={() => setCreateModalOpen(true)}
          >
            新建时段
          </Button>
        </Group>
      </Group>

      <Paper shadow="sm" withBorder>
        <table style={{ width: '100%', borderCollapse: 'collapse' }}>
          <thead>
            <tr style={{ backgroundColor: '#f8f9fa' }}>
              <th style={{ padding: '12px', textAlign: 'left' }}>时段编号</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>时段名称</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>类型</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>时间范围</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>昼间限值</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>夜间限值</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>状态</th>
            </tr>
          </thead>
          <tbody>
            {timePeriods.map((period) => (
              <tr key={period.id}>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{period.period_no}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{period.period_name}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Badge size="sm">{period.period_type}</Badge>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">
                    {period.time_from} - {period.time_to}
                  </Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{period.day_limit} dB</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{period.night_limit} dB</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Badge size="sm" color={period.status === 'active' ? 'green' : 'gray'}>
                    {period.status === 'active' ? '启用' : '禁用'}
                  </Badge>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </Paper>

      <Modal
        opened={createModalOpen}
        onClose={() => setCreateModalOpen(false)}
        title="新建时段分类"
        size="lg"
      >
        <Stack gap="md">
          <TextInput
            label="时段编号"
            required
            value={formData.period_no || ''}
            onChange={(e) => setFormData({ ...formData, period_no: e.target.value })}
          />
          <TextInput
            label="时段名称"
            required
            value={formData.period_name || ''}
            onChange={(e) => setFormData({ ...formData, period_name: e.target.value })}
          />
          <Select
            label="时段类型"
            required
            data={[
              { value: 'day', label: '昼间' },
              { value: 'night', label: '夜间' },
              { value: 'all', label: '全天' },
            ]}
            value={formData.period_type || null}
            onChange={(value) => setFormData({ ...formData, period_type: value || undefined })}
          />
          <Group grow>
            <TextInput
              label="起始时间"
              type="time"
              value={formData.time_from || ''}
              onChange={(e) => setFormData({ ...formData, time_from: e.target.value })}
            />
            <TextInput
              label="结束时间"
              type="time"
              value={formData.time_to || ''}
              onChange={(e) => setFormData({ ...formData, time_to: e.target.value })}
            />
          </Group>
          <Grid>
            <Grid.Col span={6}>
              <NumberInput
                label="昼间限值 (dB)"
                decimalScale={1}
                value={formData.day_limit}
                onChange={(value) => setFormData({ ...formData, day_limit: Number(value) })}
              />
            </Grid.Col>
            <Grid.Col span={6}>
              <NumberInput
                label="夜间限值 (dB)"
                decimalScale={1}
                value={formData.night_limit}
                onChange={(value) => setFormData({ ...formData, night_limit: Number(value) })}
              />
            </Grid.Col>
          </Grid>
          <Textarea
            label="描述"
            value={formData.description || ''}
            onChange={(e) => setFormData({ ...formData, description: e.target.value })}
          />
          <Group justify="flex-end">
            <Button variant="subtle" onClick={() => setCreateModalOpen(false)}>
              取消
            </Button>
            <Button onClick={handleCreate}>创建</Button>
          </Group>
        </Stack>
      </Modal>
    </Stack>
  );
}

import { Grid } from '@mantine/core';

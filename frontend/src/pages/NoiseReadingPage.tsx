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
  FileInput,
  Progress,
} from '@mantine/core';
import { IconPlus, IconUpload, IconFileText } from '@tabler/icons-react';
import { useNoiseReadingStore } from '../stores/noiseReadingStore';
import { NoiseReading } from '../api';

export function NoiseReadingPage() {
  const { readings, loading, fetchReadings, createReading } = useNoiseReadingStore();
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [formData, setFormData] = useState<Partial<NoiseReading>>({});

  useEffect(() => {
    fetchReadings();
  }, [fetchReadings]);

  const handleCreate = async () => {
    await createReading(formData);
    setCreateModalOpen(false);
    setFormData({});
  };

  const exceededCount = readings.filter((r) => r.is_exceeded).length;
  const complianceRate =
    readings.length > 0 ? ((readings.length - exceededCount) / readings.length) * 100 : 0;

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Text size="xl" fw={700}>
          噪声读数管理
        </Text>
        <Button
          leftSection={<IconPlus size={16} />}
          onClick={() => setCreateModalOpen(true)}
        >
          录入读数
        </Button>
      </Group>

      <Grid>
        <Grid.Col span={4}>
          <Paper shadow="sm" p="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed">
                总读数
              </Text>
              <Text size="xl" fw={700}>
                {readings.length}
              </Text>
            </Stack>
          </Paper>
        </Grid.Col>
        <Grid.Col span={4}>
          <Paper shadow="sm" p="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed">
                超标读数
              </Text>
              <Text size="xl" fw={700} c="red">
                {exceededCount}
              </Text>
            </Stack>
          </Paper>
        </Grid.Col>
        <Grid.Col span={4}>
          <Paper shadow="sm" p="md" withBorder>
            <Stack gap="xs">
              <Text size="sm" c="dimmed">
                达标率
              </Text>
              <Text size="xl" fw={700}>
                {complianceRate.toFixed(1)}%
              </Text>
            </Stack>
          </Paper>
        </Grid.Col>
      </Grid>

      <Paper shadow="sm" withBorder>
        <table style={{ width: '100%', borderCollapse: 'collapse' }}>
          <thead>
            <tr style={{ backgroundColor: '#f8f9fa' }}>
              <th style={{ padding: '12px', textAlign: 'left' }}>读数编号</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>点位</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>测量时间</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>Leq (dB)</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>标准限值</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>状态</th>
            </tr>
          </thead>
          <tbody>
            {readings.map((reading) => (
              <tr key={reading.id}>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{reading.reading_no}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{reading.point_no}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">
                    {reading.measurement_date} {reading.measurement_time}
                  </Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm" fw={500} c={reading.is_exceeded ? 'red' : undefined}>
                    {reading.leq.toFixed(1)}
                  </Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{reading.standard_limit}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Badge size="sm" color={reading.is_exceeded ? 'red' : 'green'}>
                    {reading.is_exceeded ? '超标' : '达标'}
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
        title="录入噪声读数"
        size="lg"
      >
        <Stack gap="md">
          <TextInput
            label="读数编号"
            required
            value={formData.reading_no || ''}
            onChange={(e) => setFormData({ ...formData, reading_no: e.target.value })}
          />
          <TextInput
            label="点位编号"
            required
            value={formData.point_no || ''}
            onChange={(e) => setFormData({ ...formData, point_no: e.target.value })}
          />
          <Group grow>
            <TextInput
              label="测量日期"
              type="date"
              value={formData.measurement_date || ''}
              onChange={(e) => setFormData({ ...formData, measurement_date: e.target.value })}
            />
            <TextInput
              label="测量时间"
              type="time"
              value={formData.measurement_time || ''}
              onChange={(e) => setFormData({ ...formData, measurement_time: e.target.value })}
            />
          </Group>
          <Grid>
            <Grid.Col span={6}>
              <NumberInput
                label="等效声级 Leq (dB)"
                required
                decimalScale={1}
                value={formData.leq}
                onChange={(value) => setFormData({ ...formData, leq: Number(value) })}
              />
            </Grid.Col>
            <Grid.Col span={6}>
              <NumberInput
                label="标准限值 (dB)"
                required
                decimalScale={1}
                value={formData.standard_limit}
                onChange={(value) => setFormData({ ...formData, standard_limit: Number(value) })}
              />
            </Grid.Col>
          </Grid>
          <Grid>
            <Grid.Col span={4}>
              <NumberInput
                label="Lmax (dB)"
                decimalScale={1}
                value={formData.lmax}
                onChange={(value) => setFormData({ ...formData, lmax: Number(value) })}
              />
            </Grid.Col>
            <Grid.Col span={4}>
              <NumberInput
                label="Lmin (dB)"
                decimalScale={1}
                value={formData.lmin}
                onChange={(value) => setFormData({ ...formData, lmin: Number(value) })}
              />
            </Grid.Col>
            <Grid.Col span={4}>
              <NumberInput
                label="L10 (dB)"
                decimalScale={1}
                value={formData.l10}
                onChange={(value) => setFormData({ ...formData, l10: Number(value) })}
              />
            </Grid.Col>
          </Grid>
          <Group justify="flex-end">
            <Button variant="subtle" onClick={() => setCreateModalOpen(false)}>
              取消
            </Button>
            <Button onClick={handleCreate}>保存</Button>
          </Group>
        </Stack>
      </Modal>
    </Stack>
  );
}

import { Grid } from '@mantine/core';

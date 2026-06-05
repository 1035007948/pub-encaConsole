import { useState } from 'react';
import {
  Stack,
  Paper,
  Group,
  Button,
  Text,
  Badge,
  Select,
  DatePickerInput,
  Grid,
} from '@mantine/core';
import { IconDownload, IconArchive } from '@tabler/icons-react';
import { ChartWidget } from '../components/ChartWidget';

export function ArchiveSnapshotPage() {
  const [selectedBatch, setSelectedBatch] = useState<string | null>(null);
  const [selectedDate, setSelectedDate] = useState<Date | null>(null);

  const mockSnapshots = [
    {
      id: 1,
      snapshot_no: 'ARCH-2024-001',
      batch_no: 'BATCH-2024-001',
      created_at: '2024-03-15 16:00',
      complaint_count: 25,
      reading_count: 150,
      evidence_count: 80,
      status: 'completed',
    },
    {
      id: 2,
      snapshot_no: 'ARCH-2024-002',
      batch_no: 'BATCH-2024-002',
      created_at: '2024-03-20 16:00',
      complaint_count: 30,
      reading_count: 180,
      evidence_count: 95,
      status: 'completed',
    },
  ];

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Text size="xl" fw={700}>
          归档快照与导出
        </Text>
        <Group>
          <Button
            variant="light"
            leftSection={<IconArchive size={16} />}
          >
            创建快照
          </Button>
          <Button
            leftSection={<IconDownload size={16} />}
          >
            导出报告
          </Button>
        </Group>
      </Group>

      <Grid>
        <Grid.Col span={4}>
          <Select
            label="选择批次"
            data={[
              { value: 'BATCH-2024-001', label: 'BATCH-2024-001' },
              { value: 'BATCH-2024-002', label: 'BATCH-2024-002' },
              { value: 'BATCH-2024-003', label: 'BATCH-2024-003' },
            ]}
            value={selectedBatch}
            onChange={setSelectedBatch}
            placeholder="选择批次"
          />
        </Grid.Col>
        <Grid.Col span={4}>
          <DatePickerInput
            label="选择日期"
            placeholder="选择日期"
            value={selectedDate}
            onChange={setSelectedDate}
            clearable
          />
        </Grid.Col>
        <Grid.Col span={4}>
          <Button fullWidth mt="xl">
            查询快照
          </Button>
        </Grid.Col>
      </Grid>

      <Paper shadow="sm" withBorder>
        <table style={{ width: '100%', borderCollapse: 'collapse' }}>
          <thead>
            <tr style={{ backgroundColor: '#f8f9fa' }}>
              <th style={{ padding: '12px', textAlign: 'left' }}>快照编号</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>批次</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>创建时间</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>投诉单数</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>读数数</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>证据数</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>状态</th>
              <th style={{ padding: '12px', textAlign: 'left' }}>操作</th>
            </tr>
          </thead>
          <tbody>
            {mockSnapshots.map((snapshot) => (
              <tr key={snapshot.id}>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{snapshot.snapshot_no}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{snapshot.batch_no}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{snapshot.created_at}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{snapshot.complaint_count}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{snapshot.reading_count}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Text size="sm">{snapshot.evidence_count}</Text>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Badge color="green">已完成</Badge>
                </td>
                <td style={{ padding: '12px', borderBottom: '1px solid #dee2e6' }}>
                  <Button size="xs" variant="light">
                    下载
                  </Button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </Paper>

      <Grid>
        <Grid.Col span={6}>
          <ChartWidget
            title="归档趋势"
            type="bar"
            data={[
              { label: '第1周', value: 25 },
              { label: '第2周', value: 30 },
              { label: '第3周', value: 28 },
              { label: '第4周', value: 35 },
            ]}
            height={200}
          />
        </Grid.Col>
        <Grid.Col span={6}>
          <ChartWidget
            title="证据完整度"
            type="pie"
            data={[
              { label: '完整', value: 85, color: '#40c057' },
              { label: '待补充', value: 15, color: '#fab005' },
            ]}
            height={200}
          />
        </Grid.Col>
      </Grid>
    </Stack>
  );
}

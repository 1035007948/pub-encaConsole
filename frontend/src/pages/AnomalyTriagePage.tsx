import { useEffect, useState } from 'react';
import {
  Stack,
  Paper,
  Group,
  Text,
  Badge,
  Modal,
  Textarea,
  Button,
  Select,
} from '@mantine/core';
import { IconAlertTriangle, IconCheck, IconClock } from '@tabler/icons-react';
import { useAnomalyStore } from '../stores/anomalyStore';
import { AnomalyEvent } from '../api';

export function AnomalyTriagePage() {
  const { anomalies, loading, fetchAnomalies, resolveAnomaly } = useAnomalyStore();
  const [resolveModalOpen, setResolveModalOpen] = useState(false);
  const [selectedAnomaly, setSelectedAnomaly] = useState<AnomalyEvent | null>(null);
  const [resolutionNote, setResolutionNote] = useState('');
  const [statusFilter, setStatusFilter] = useState<string | null>('open');

  useEffect(() => {
    const params: Record<string, string> = {};
    if (statusFilter) {
      params['status'] = statusFilter;
    }
    fetchAnomalies(params);
  }, [fetchAnomalies, statusFilter]);

  const handleResolve = async () => {
    if (selectedAnomaly && resolutionNote) {
      await resolveAnomaly(selectedAnomaly.id, resolutionNote);
      setResolveModalOpen(false);
      setSelectedAnomaly(null);
      setResolutionNote('');
    }
  };

  const getSeverityColor = (severity: string) => {
    const colorMap: Record<string, string> = {
      critical: 'red',
      high: 'orange',
      medium: 'yellow',
      low: 'blue',
    };
    return colorMap[severity] || 'gray';
  };

  const getStatusColor = (status: string) => {
    const colorMap: Record<string, string> = {
      open: 'red',
      in_progress: 'yellow',
      resolved: 'green',
    };
    return colorMap[status] || 'gray';
  };

  const getStatusLabel = (status: string) => {
    const labelMap: Record<string, string> = {
      open: '待处理',
      in_progress: '处理中',
      resolved: '已解决',
    };
    return labelMap[status] || status;
  };

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Text size="xl" fw={700}>
          异常分诊
        </Text>
        <Select
          data={[
            { value: 'open', label: '待处理' },
            { value: 'in_progress', label: '处理中' },
            { value: 'resolved', label: '已解决' },
            { value: '', label: '全部' },
          ]}
          value={statusFilter}
          onChange={setStatusFilter}
          placeholder="筛选状态"
          clearable
        />
      </Group>

      <Paper shadow="sm" withBorder>
        <Stack gap={0}>
          {anomalies.map((anomaly, index) => (
            <Paper
              key={anomaly.id}
              p="md"
              style={{
                borderBottom:
                  index < anomalies.length - 1 ? '1px solid #e9ecef' : undefined,
                backgroundColor: anomaly.status === 'open' ? '#fff5f5' : undefined,
              }}
            >
              <Stack gap="sm">
                <Group justify="space-between">
                  <Group>
                    <IconAlertTriangle
                      size={20}
                      color={getSeverityColor(anomaly.severity)}
                    />
                    <Text fw={500}>{anomaly.event_no}</Text>
                    <Badge size="sm" color={getSeverityColor(anomaly.severity)}>
                      {anomaly.severity}
                    </Badge>
                  </Group>
                  <Badge color={getStatusColor(anomaly.status)}>
                    {getStatusLabel(anomaly.status)}
                  </Badge>
                </Group>

                <Text size="sm">{anomaly.event_name}</Text>

                <Group gap="md">
                  <Text size="xs" c="dimmed">
                    类型: {anomaly.event_type}
                  </Text>
                  <Text size="xs" c="dimmed">
                    实体: {anomaly.entity_type} #{anomaly.entity_no}
                  </Text>
                  <Text size="xs" c="dimmed">
                    触发字段: {anomaly.trigger_field}
                  </Text>
                </Group>

                <Group gap="md">
                  <Text size="xs" c="dimmed">
                    触发值: {anomaly.trigger_value}
                  </Text>
                  <Text size="xs" c="dimmed">
                    阈值: {anomaly.threshold_value}
                  </Text>
                </Group>

                {anomaly.status === 'open' && (
                  <Group>
                    <Button
                      size="xs"
                      variant="light"
                      onClick={() => {
                        setSelectedAnomaly(anomaly);
                        setResolveModalOpen(true);
                      }}
                    >
                      处理异常
                    </Button>
                  </Group>
                )}

                {anomaly.status === 'resolved' && anomaly.resolved_at && (
                  <Group gap="xs">
                    <IconCheck size={14} color="green" />
                    <Text size="xs" c="dimmed">
                      已于 {new Date(anomaly.resolved_at).toLocaleString()} 解决
                    </Text>
                    {anomaly.resolution_note && (
                      <Text size="xs" c="dimmed">
                        - {anomaly.resolution_note}
                      </Text>
                    )}
                  </Group>
                )}
              </Stack>
            </Paper>
          ))}

          {anomalies.length === 0 && (
            <Group justify="center" p="xl">
              <IconCheck size={24} color="green" />
              <Text c="dimmed">暂无异常事件</Text>
            </Group>
          )}
        </Stack>
      </Paper>

      <Modal
        opened={resolveModalOpen}
        onClose={() => setResolveModalOpen(false)}
        title="处理异常"
      >
        <Stack gap="md">
          {selectedAnomaly && (
            <>
              <Text size="sm" c="dimmed">
                异常编号: {selectedAnomaly.event_no}
              </Text>
              <Text size="sm">{selectedAnomaly.event_name}</Text>
              <Text size="sm" c="dimmed">
                触发字段: {selectedAnomaly.trigger_field} = {selectedAnomaly.trigger_value}
                (阈值: {selectedAnomaly.threshold_value})
              </Text>
            </>
          )}
          <Textarea
            label="处理说明"
            placeholder="请输入处理说明..."
            required
            value={resolutionNote}
            onChange={(e) => setResolutionNote(e.target.value)}
            minRows={3}
          />
          <Group justify="flex-end">
            <Button variant="subtle" onClick={() => setResolveModalOpen(false)}>
              取消
            </Button>
            <Button onClick={handleResolve} disabled={!resolutionNote.trim()}>
              确认解决
            </Button>
          </Group>
        </Stack>
      </Modal>
    </Stack>
  );
}

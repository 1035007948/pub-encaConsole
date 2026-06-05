import { useEffect, useState } from 'react';
import {
  Stack,
  Paper,
  Group,
  Button,
  Text,
  Badge,
  Grid,
  Modal,
  TextInput,
  Select,
  Textarea,
  NumberInput,
} from '@mantine/core';
import { IconPlus, IconRefresh, IconSettings } from '@tabler/icons-react';

interface RuleConfig {
  id: number;
  rule_no: string;
  rule_name: string;
  rule_type: string;
  description: string;
  condition: string;
  action: string;
  priority: number;
  status: string;
  created_at: string;
  updated_at: string;
}

export function RuleConfigPage() {
  const [rules, setRules] = useState<RuleConfig[]>([]);
  const [createModalOpen, setCreateModalOpen] = useState(false);

  useEffect(() => {
    fetchRules();
  }, []);

  const fetchRules = async () => {
    setRules([
      {
        id: 1,
        rule_no: 'RULE-001',
        rule_name: '时段不合规检测',
        rule_type: 'compliance',
        description: '检测采样时段是否符合规定',
        condition: 'measurement_time NOT IN (day_period OR night_period)',
        action: 'create_anomaly',
        priority: 1,
        status: 'active',
        created_at: '2024-01-01',
        updated_at: '2024-01-01',
      },
      {
        id: 2,
        rule_no: 'RULE-002',
        rule_name: '超标读数检测',
        rule_type: 'threshold',
        description: '检测噪声读数是否超标',
        condition: 'leq > standard_limit',
        action: 'mark_exceeded',
        priority: 2,
        status: 'active',
        created_at: '2024-01-01',
        updated_at: '2024-01-01',
      },
      {
        id: 3,
        rule_no: 'RULE-003',
        rule_name: '证据完整度检测',
        rule_type: 'completeness',
        description: '检测证据附件是否完整',
        condition: 'evidence_count < required_count',
        action: 'require_supplement',
        priority: 3,
        status: 'active',
        created_at: '2024-01-01',
        updated_at: '2024-01-01',
      },
    ]);
  };

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Text size="xl" fw={700}>
          规则配置
        </Text>
        <Group>
          <Button
            variant="light"
            leftSection={<IconRefresh size={16} />}
            onClick={fetchRules}
          >
            刷新
          </Button>
          <Button
            leftSection={<IconPlus size={16} />}
            onClick={() => setCreateModalOpen(true)}
          >
            新建规则
          </Button>
        </Group>
      </Group>

      <Grid>
        {rules.map((rule) => (
          <Grid.Col key={rule.id} span={6}>
            <Paper shadow="sm" p="md" withBorder>
              <Stack gap="sm">
                <Group justify="space-between">
                  <Text fw={500}>{rule.rule_no}</Text>
                  <Badge color={rule.status === 'active' ? 'green' : 'gray'}>
                    {rule.status === 'active' ? '启用' : '禁用'}
                  </Badge>
                </Group>
                <Text size="sm">{rule.rule_name}</Text>
                <Text size="xs" c="dimmed">
                  {rule.description}
                </Text>
                <Stack gap="xs">
                  <Text size="xs" c="dimmed">
                    条件: {rule.condition}
                  </Text>
                  <Text size="xs" c="dimmed">
                    动作: {rule.action}
                  </Text>
                </Stack>
                <Group>
                  <Badge size="sm">{rule.rule_type}</Badge>
                  <Text size="xs" c="dimmed">
                    优先级: {rule.priority}
                  </Text>
                </Group>
              </Stack>
            </Paper>
          </Grid.Col>
        ))}
      </Grid>

      <Modal
        opened={createModalOpen}
        onClose={() => setCreateModalOpen(false)}
        title="新建规则"
        size="lg"
      >
        <Stack gap="md">
          <TextInput label="规则编号" required />
          <TextInput label="规则名称" required />
          <Select
            label="规则类型"
            required
            data={[
              { value: 'compliance', label: '合规检测' },
              { value: 'threshold', label: '阈值检测' },
              { value: 'completeness', label: '完整度检测' },
            ]}
          />
          <Textarea label="描述" />
          <Textarea label="条件表达式" placeholder="例如: leq > standard_limit" />
          <Select
            label="触发动作"
            data={[
              { value: 'create_anomaly', label: '创建异常事件' },
              { value: 'mark_exceeded', label: '标记超标' },
              { value: 'require_supplement', label: '要求补充' },
            ]}
          />
          <NumberInput label="优先级" min={1} max={10} />
          <Group justify="flex-end">
            <Button variant="subtle" onClick={() => setCreateModalOpen(false)}>
              取消
            </Button>
            <Button onClick={() => setCreateModalOpen(false)}>创建</Button>
          </Group>
        </Stack>
      </Modal>
    </Stack>
  );
}

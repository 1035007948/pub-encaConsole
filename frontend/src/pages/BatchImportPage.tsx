import { useState } from 'react';
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
  FileInput,
  Progress,
} from '@mantine/core';
import { IconUpload, IconFileSpreadsheet, IconCheck, IconX } from '@tabler/icons-react';
import { ImportPreviewTable } from '../components/ImportPreviewTable';

export function BatchImportPage() {
  const [importModalOpen, setImportModalOpen] = useState(false);
  const [importing, setImporting] = useState(false);
  const [importProgress, setImportProgress] = useState(0);

  const mockColumns = [
    { source: '时段编号', target: 'period_no', required: true, valid: true },
    { source: '时段名称', target: 'period_name', required: true, valid: true },
    { source: '时段类型', target: 'period_type', required: true, valid: true },
    { source: '起始时间', target: 'time_from', required: true, valid: true },
    { source: '结束时间', target: 'time_to', required: true, valid: true },
    { source: '昼间限值', target: 'day_limit', required: false, valid: true },
    { source: '夜间限值', target: 'night_limit', required: false, valid: true },
  ];

  const mockRows = [
    {
      index: 0,
      data: {
        '时段编号': 'TP-2024-0001',
        '时段名称': '昼间时段',
        '时段类型': 'day',
        '起始时间': '06:00',
        '结束时间': '22:00',
        '昼间限值': '65',
        '夜间限值': '55',
      },
      valid: true,
    },
    {
      index: 1,
      data: {
        '时段编号': 'TP-2024-0002',
        '时段名称': '夜间时段',
        '时段类型': 'night',
        '起始时间': '22:00',
        '结束时间': '06:00',
        '昼间限值': '55',
        '夜间限值': '45',
      },
      valid: true,
    },
    {
      index: 2,
      data: {
        '时段编号': 'TP-2024-0003',
        '时段名称': '特殊时段',
        '时段类型': 'special',
        '起始时间': '12:00',
        '结束时间': '14:00',
        '昼间限值': '',
        '夜间限值': '',
      },
      valid: false,
      errors: { '时段类型': '无效的时段类型' },
    },
  ];

  const handleImport = () => {
    setImporting(true);
    setImportProgress(0);

    const interval = setInterval(() => {
      setImportProgress((prev) => {
        if (prev >= 100) {
          clearInterval(interval);
          setImporting(false);
          return 100;
        }
        return prev + 10;
      });
    }, 200);
  };

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Text size="xl" fw={700}>
          时段分类批量导入
        </Text>
        <Button
          leftSection={<IconUpload size={16} />}
          onClick={() => setImportModalOpen(true)}
        >
          导入数据
        </Button>
      </Group>

      <Paper shadow="sm" p="md" withBorder>
        <Stack gap="md">
          <Text fw={500}>导入说明</Text>
          <Text size="sm" c="dimmed">
            支持CSV、Excel格式的时段分类数据批量导入。导入前请确保数据格式正确，系统将自动进行数据校验。
          </Text>
          <Group>
            <Button variant="light" leftSection={<IconFileSpreadsheet size={16} />}>
              下载模板
            </Button>
          </Group>
        </Stack>
      </Paper>

      <Paper shadow="sm" p="md" withBorder>
        <Stack gap="md">
          <Text fw={500}>导入历史</Text>
          <Stack gap="xs">
            <Group justify="space-between" p="sm" style={{ backgroundColor: '#f8f9fa' }}>
              <Group>
                <IconCheck size={16} color="green" />
                <Text size="sm">时段分类导入 - 2024-03-15 14:30</Text>
              </Group>
              <Badge color="green">成功 50 条</Badge>
            </Group>
            <Group justify="space-between" p="sm" style={{ backgroundColor: '#f8f9fa' }}>
              <Group>
                <IconX size={16} color="red" />
                <Text size="sm">时段分类导入 - 2024-03-14 10:15</Text>
              </Group>
              <Badge color="red">失败 3 条</Badge>
            </Group>
          </Stack>
        </Stack>
      </Paper>

      <Modal
        opened={importModalOpen}
        onClose={() => setImportModalOpen(false)}
        title="批量导入时段分类"
        size="xl"
      >
        <Stack gap="md">
          {!importing ? (
            <>
              <FileInput
                label="选择文件"
                placeholder="点击或拖拽文件到此区域"
                accept=".csv,.xlsx,.xls"
              />

              <ImportPreviewTable
                columns={mockColumns}
                rows={mockRows}
                onColumnMap={(source, target) => {
                  console.log('Map column:', source, '->', target);
                }}
                onConfirm={handleImport}
                onCancel={() => setImportModalOpen(false)}
              />
            </>
          ) : (
            <Stack gap="md" align="center" py="xl">
              <Text>正在导入数据...</Text>
              <Progress value={importProgress} size="lg" w="100%" />
              <Text size="sm" c="dimmed">
                {importProgress}% 完成
              </Text>
            </Stack>
          )}
        </Stack>
      </Modal>
    </Stack>
  );
}

import { useState } from 'react';
import {
  Paper,
  Stack,
  Group,
  Text,
  Badge,
  Table,
  Select,
  Button,
  ScrollArea,
  Alert,
} from '@mantine/core';
import { IconAlertCircle, IconCheck, IconX } from '@tabler/icons-react';

interface ImportColumn {
  source: string;
  target?: string;
  required?: boolean;
  valid?: boolean;
  error?: string;
}

interface ImportRow {
  index: number;
  data: Record<string, string>;
  valid?: boolean;
  errors?: Record<string, string>;
}

interface ImportPreviewTableProps {
  columns: ImportColumn[];
  rows: ImportRow[];
  onColumnMap: (source: string, target: string) => void;
  onConfirm: () => void;
  onCancel: () => void;
}

export function ImportPreviewTable({
  columns,
  rows,
  onColumnMap,
  onConfirm,
  onCancel,
}: ImportPreviewTableProps) {
  const [currentPage, setCurrentPage] = useState(0);
  const pageSize = 20;
  const totalPages = Math.ceil(rows.length / pageSize);

  const validRows = rows.filter((r) => r.valid !== false).length;
  const invalidRows = rows.length - validRows;

  const targetOptions = [
    { value: 'period_no', label: '时段编号' },
    { value: 'period_name', label: '时段名称' },
    { value: 'period_type', label: '时段类型' },
    { value: 'time_from', label: '起始时间' },
    { value: 'time_to', label: '结束时间' },
    { value: 'day_limit', label: '昼间限值' },
    { value: 'night_limit', label: '夜间限值' },
    { value: 'description', label: '描述' },
  ];

  const paginatedRows = rows.slice(
    currentPage * pageSize,
    (currentPage + 1) * pageSize
  );

  return (
    <Paper shadow="sm" withBorder>
      <Stack gap="md">
        <Group justify="space-between" p="md">
          <Text fw={500}>导入预览</Text>
          <Group>
            <Badge color="green" leftSection={<IconCheck size={12} />}>
              {validRows} 有效
            </Badge>
            {invalidRows > 0 && (
              <Badge color="red" leftSection={<IconX size={12} />}>
                {invalidRows} 无效
              </Badge>
            )}
          </Group>
        </Group>

        <Stack px="md">
          <Text size="sm" fw={500}>
            字段映射
          </Text>
          <Group gap="md">
            {columns.map((col) => (
              <Group key={col.source} gap="xs">
                <Text size="sm">{col.source}</Text>
                <Text size="sm" c="dimmed">
                  →
                </Text>
                <Select
                  size="xs"
                  data={targetOptions}
                  value={col.target || null}
                  onChange={(value) => value && onColumnMap(col.source, value)}
                  placeholder="选择目标字段"
                  clearable
                />
                {col.required && (
                  <Badge size="xs" color="red">
                    必填
                  </Badge>
                )}
                {col.error && (
                  <Text size="xs" c="red">
                    {col.error}
                  </Text>
                )}
              </Group>
            ))}
          </Group>
        </Stack>

        {invalidRows > 0 && (
          <Alert
            icon={<IconAlertCircle size={16} />}
            title="存在无效数据"
            color="yellow"
            mx="md"
          >
            共 {invalidRows} 行数据存在问题，请检查数据格式是否正确
          </Alert>
        )}

        <ScrollArea h={300}>
          <Table striped highlightOnHover withTableBorder>
            <Table.Thead>
              <Table.Tr>
                <Table.Th style={{ width: 50 }}>#</Table.Th>
                {columns.map((col) => (
                  <Table.Th key={col.source}>
                    {col.target || col.source}
                    {col.required && (
                      <Badge size="xs" color="red" ml="xs">
                        *
                      </Badge>
                    )}
                  </Table.Th>
                ))}
                <Table.Th style={{ width: 80 }}>状态</Table.Th>
              </Table.Tr>
            </Table.Thead>
            <Table.Tbody>
              {paginatedRows.map((row) => (
                <Table.Tr key={row.index} style={{ opacity: row.valid === false ? 0.5 : 1 }}>
                  <Table.Td>{row.index + 1}</Table.Td>
                  {columns.map((col) => (
                    <Table.Td key={col.source}>
                      <Text
                        size="sm"
                        c={row.errors?.[col.source || ''] ? 'red' : undefined}
                      >
                        {row.data[col.source] || '-'}
                      </Text>
                    </Table.Td>
                  ))}
                  <Table.Td>
                    {row.valid === false ? (
                      <Badge size="sm" color="red">
                        无效
                      </Badge>
                    ) : (
                      <Badge size="sm" color="green">
                        有效
                      </Badge>
                    )}
                  </Table.Td>
                </Table.Tr>
              ))}
            </Table.Tbody>
          </Table>
        </ScrollArea>

        <Group justify="space-between" px="md" pb="md">
          <Group>
            <Text size="sm" c="dimmed">
              第 {currentPage + 1} / {totalPages} 页
            </Text>
          </Group>
          <Group>
            <Button variant="subtle" onClick={onCancel}>
              取消
            </Button>
            <Button
              onClick={onConfirm}
              disabled={validRows === 0 || columns.some((c) => c.required && !c.target)}
            >
              确认导入 ({validRows} 条)
            </Button>
          </Group>
        </Group>
      </Stack>
    </Paper>
  );
}

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
  Textarea,
  NumberInput,
} from '@mantine/core';
import { IconPlus, IconRefresh } from '@tabler/icons-react';
import { useComplaintStore } from '../stores/complaintStore';
import { AdvancedFilterBar } from '../components/AdvancedFilterBar';
import { BatchActionToolbar } from '../components/BatchActionToolbar';
import { DetailDrawer } from '../components/DetailDrawer';
import { Complaint } from '../api';
import {
  createColumnHelper,
  flexRender,
  getCoreRowModel,
  useReactTable,
} from '@tanstack/react-table';

const columnHelper = createColumnHelper<Complaint>();

export function ComplaintListPage() {
  const {
    complaints,
    total,
    loading,
    fetchComplaints,
    createComplaint,
    deleteComplaint,
    transitionComplaint,
    currentComplaint,
    setCurrentComplaint,
  } = useComplaintStore();

  const [filters, setFilters] = useState<Record<string, unknown>>({});
  const [selectedIds, setSelectedIds] = useState<number[]>([]);
  const [createModalOpen, setCreateModalOpen] = useState(false);
  const [formData, setFormData] = useState<Partial<Complaint>>({});

  useEffect(() => {
    fetchComplaints();
  }, [fetchComplaints]);

  const columns = [
    columnHelper.accessor('complaint_no', {
      header: '投诉编号',
      cell: (info) => (
        <Text
          size="sm"
          style={{ cursor: 'pointer' }}
          onClick={() => setCurrentComplaint(info.row.original)}
        >
          {info.getValue()}
        </Text>
      ),
    }),
    columnHelper.accessor('title', {
      header: '标题',
      cell: (info) => <Text size="sm">{info.getValue()}</Text>,
    }),
    columnHelper.accessor('status', {
      header: '状态',
      cell: (info) => {
        const status = info.getValue();
        const colorMap: Record<string, string> = {
          draft: 'gray',
          pending: 'yellow',
          in_progress: 'blue',
          completed: 'green',
          rejected: 'red',
        };
        const labelMap: Record<string, string> = {
          draft: '草稿',
          pending: '待处理',
          in_progress: '进行中',
          completed: '已完成',
          rejected: '已驳回',
        };
        return (
          <Badge size="sm" color={colorMap[status] || 'gray'}>
            {labelMap[status] || status}
          </Badge>
        );
      },
    }),
    columnHelper.accessor('level', {
      header: '等级',
      cell: (info) => {
        const level = info.getValue();
        const colorMap: Record<string, string> = {
          urgent: 'red',
          high: 'orange',
          medium: 'yellow',
          low: 'green',
        };
        return (
          <Badge size="sm" color={colorMap[level] || 'gray'}>
            {level}
          </Badge>
        );
      },
    }),
    columnHelper.accessor('enterprise_name', {
      header: '企业名称',
      cell: (info) => <Text size="sm">{info.getValue()}</Text>,
    }),
    columnHelper.accessor('responsible_user', {
      header: '负责人',
      cell: (info) => <Text size="sm">{info.getValue()}</Text>,
    }),
    columnHelper.accessor('created_at', {
      header: '创建时间',
      cell: (info) => (
        <Text size="sm" c="dimmed">
          {new Date(info.getValue()).toLocaleDateString()}
        </Text>
      ),
    }),
  ];

  const table = useReactTable({
    data: complaints,
    columns,
    getCoreRowModel: getCoreRowModel(),
  });

  const filterFields = [
    { name: 'status', label: '状态', type: 'select' as const, options: [
      { value: 'draft', label: '草稿' },
      { value: 'pending', label: '待处理' },
      { value: 'in_progress', label: '进行中' },
      { value: 'completed', label: '已完成' },
      { value: 'rejected', label: '已驳回' },
    ]},
    { name: 'level', label: '等级', type: 'select' as const, options: [
      { value: 'urgent', label: '紧急' },
      { value: 'high', label: '高' },
      { value: 'medium', label: '中' },
      { value: 'low', label: '低' },
    ]},
    { name: 'enterprise_name', label: '企业名称', type: 'text' as const },
  ];

  const batchActions = [
    {
      key: 'delete',
      label: '删除',
      color: 'red',
      onExecute: async (ids: number[]) => {
        for (const id of ids) {
          await deleteComplaint(id);
        }
        setSelectedIds([]);
      },
    },
    {
      key: 'complete',
      label: '标记完成',
      color: 'green',
      onExecute: async (ids: number[]) => {
        for (const id of ids) {
          await transitionComplaint(id, 'completed');
        }
        setSelectedIds([]);
      },
    },
  ];

  const handleCreate = async () => {
    await createComplaint(formData);
    setCreateModalOpen(false);
    setFormData({});
  };

  return (
    <Stack gap="md">
      <Group justify="space-between">
        <Text size="xl" fw={700}>
          投诉单管理
        </Text>
        <Group>
          <Button
            variant="light"
            leftSection={<IconRefresh size={16} />}
            onClick={() => fetchComplaints()}
          >
            刷新
          </Button>
          <Button
            leftSection={<IconPlus size={16} />}
            onClick={() => setCreateModalOpen(true)}
          >
            新建投诉单
          </Button>
        </Group>
      </Group>

      <AdvancedFilterBar
        fields={filterFields}
        values={filters}
        onChange={setFilters}
        onReset={() => fetchComplaints()}
      />

      <BatchActionToolbar
        items={complaints}
        selectedIds={selectedIds}
        onSelectionChange={setSelectedIds}
        getId={(c) => c.id}
        actions={batchActions}
      />

      <Paper shadow="sm" withBorder>
        <table style={{ width: '100%', borderCollapse: 'collapse' }}>
          <thead>
            <tr style={{ backgroundColor: '#f8f9fa' }}>
              {table.getFlatHeaders().map((header) => (
                <th
                  key={header.id}
                  style={{
                    padding: '12px',
                    textAlign: 'left',
                    borderBottom: '1px solid #dee2e6',
                  }}
                >
                  {flexRender(
                    header.column.columnDef.header,
                    header.getContext()
                  )}
                </th>
              ))}
            </tr>
          </thead>
          <tbody>
            {table.getRowModel().rows.map((row) => (
              <tr key={row.id}>
                {row.getVisibleCells().map((cell) => (
                  <td
                    key={cell.id}
                    style={{
                      padding: '12px',
                      borderBottom: '1px solid #dee2e6',
                    }}
                  >
                    {flexRender(
                      cell.column.columnDef.cell,
                      cell.getContext()
                    )}
                  </td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </Paper>

      <Group justify="space-between">
        <Text size="sm" c="dimmed">
          共 {total} 条记录
        </Text>
      </Group>

      <DetailDrawer
        opened={currentComplaint !== null}
        onClose={() => setCurrentComplaint(null)}
        title="投诉单详情"
        data={currentComplaint}
        fields={[
          { key: 'complaint_no', label: '投诉编号' },
          { key: 'title', label: '标题' },
          { key: 'description', label: '描述' },
          { key: 'status', label: '状态', type: 'badge' },
          { key: 'level', label: '等级', type: 'badge' },
          { key: 'complainant_name', label: '投诉人' },
          { key: 'complainant_tel', label: '联系电话' },
          { key: 'enterprise_name', label: '企业名称' },
          { key: 'enterprise_addr', label: '企业地址' },
          { key: 'responsible_user', label: '负责人' },
          { key: 'priority', label: '优先级' },
          { key: 'created_at', label: '创建时间', type: 'date' },
        ]}
        actions={[
          {
            label: '编辑',
            onClick: () => {},
          },
          {
            label: '删除',
            onClick: async () => {
              if (currentComplaint) {
                await deleteComplaint(currentComplaint.id);
                setCurrentComplaint(null);
              }
            },
            color: 'red',
          },
        ]}
      />

      <Modal
        opened={createModalOpen}
        onClose={() => setCreateModalOpen(false)}
        title="新建投诉单"
        size="lg"
      >
        <Stack gap="md">
          <TextInput
            label="标题"
            required
            value={formData.title || ''}
            onChange={(e) => setFormData({ ...formData, title: e.target.value })}
          />
          <Textarea
            label="描述"
            value={formData.description || ''}
            onChange={(e) => setFormData({ ...formData, description: e.target.value })}
          />
          <Select
            label="等级"
            required
            data={[
              { value: 'urgent', label: '紧急' },
              { value: 'high', label: '高' },
              { value: 'medium', label: '中' },
              { value: 'low', label: '低' },
            ]}
            value={formData.level || null}
            onChange={(value) => setFormData({ ...formData, level: value || undefined })}
          />
          <TextInput
            label="投诉人"
            value={formData.complainant_name || ''}
            onChange={(e) => setFormData({ ...formData, complainant_name: e.target.value })}
          />
          <TextInput
            label="联系电话"
            value={formData.complainant_tel || ''}
            onChange={(e) => setFormData({ ...formData, complainant_tel: e.target.value })}
          />
          <TextInput
            label="企业名称"
            required
            value={formData.enterprise_name || ''}
            onChange={(e) => setFormData({ ...formData, enterprise_name: e.target.value })}
          />
          <TextInput
            label="企业地址"
            value={formData.enterprise_addr || ''}
            onChange={(e) => setFormData({ ...formData, enterprise_addr: e.target.value })}
          />
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
    </Stack>
  );
}

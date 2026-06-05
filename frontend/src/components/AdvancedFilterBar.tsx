import { useState } from 'react';
import {
  Paper,
  Stack,
  Group,
  TextInput,
  Select,
  DatePickerInput,
  Button,
  Collapse,
  ActionIcon,
  Badge,
} from '@mantine/core';
import { IconFilter, IconChevronDown, IconX } from '@tabler/icons-react';

interface FilterField {
  name: string;
  label: string;
  type: 'text' | 'select' | 'date' | 'dateRange';
  options?: { value: string; label: string }[];
}

interface AdvancedFilterBarProps {
  fields: FilterField[];
  values: Record<string, unknown>;
  onChange: (values: Record<string, unknown>) => void;
  onReset?: () => void;
}

export function AdvancedFilterBar({
  fields,
  values,
  onChange,
  onReset,
}: AdvancedFilterBarProps) {
  const [expanded, setExpanded] = useState(false);
  const activeFilterCount = Object.values(values).filter(
    (v) => v !== undefined && v !== null && v !== ''
  ).length;

  const handleChange = (name: string, value: unknown) => {
    onChange({ ...values, [name]: value });
  };

  const handleReset = () => {
    onChange({});
    onReset?.();
  };

  return (
    <Paper shadow="sm" p="md" withBorder>
      <Stack gap="sm">
        <Group justify="space-between">
          <Group>
            <IconFilter size={20} />
            <span>高级筛选</span>
            {activeFilterCount > 0 && (
              <Badge size="sm" variant="filled">
                {activeFilterCount} 项
              </Badge>
            )}
          </Group>
          <Group>
            {activeFilterCount > 0 && (
              <Button size="xs" variant="subtle" onClick={handleReset}>
                重置
              </Button>
            )}
            <ActionIcon
              variant="subtle"
              onClick={() => setExpanded(!expanded)}
            >
              <IconChevronDown
                size={16}
                style={{ transform: expanded ? 'rotate(180deg)' : 'none' }}
              />
            </ActionIcon>
          </Group>
        </Group>

        <Collapse in={expanded}>
          <Group gap="md" grow>
            {fields.map((field) => {
              switch (field.type) {
                case 'text':
                  return (
                    <TextInput
                      key={field.name}
                      label={field.label}
                      placeholder={`输入${field.label}`}
                      value={(values[field.name] as string) || ''}
                      onChange={(e) => handleChange(field.name, e.target.value)}
                      rightSection={
                        values[field.name] ? (
                          <ActionIcon
                            size="xs"
                            variant="subtle"
                            onClick={() => handleChange(field.name, '')}
                          >
                            <IconX size={12} />
                          </ActionIcon>
                        ) : null
                      }
                    />
                  );

                case 'select':
                  return (
                    <Select
                      key={field.name}
                      label={field.label}
                      placeholder={`选择${field.label}`}
                      data={field.options || []}
                      value={(values[field.name] as string) || null}
                      onChange={(value) => handleChange(field.name, value)}
                      clearable
                    />
                  );

                case 'date':
                  return (
                    <DatePickerInput
                      key={field.name}
                      label={field.label}
                      placeholder={`选择${field.label}`}
                      value={values[field.name] as Date}
                      onChange={(value) => handleChange(field.name, value)}
                      clearable
                    />
                  );

                default:
                  return null;
              }
            })}
          </Group>
        </Collapse>
      </Stack>
    </Paper>
  );
}

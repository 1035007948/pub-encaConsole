import { Paper, Text, Group, Stack, Box } from '@mantine/core';

interface ChartData {
  label: string;
  value: number;
  color?: string;
}

interface ChartWidgetProps {
  title: string;
  type: 'bar' | 'pie' | 'line' | 'metric';
  data?: ChartData[];
  value?: number;
  unit?: string;
  trend?: number;
  height?: number;
  onDrillDown?: (label: string) => void;
}

export function ChartWidget({
  title,
  type,
  data,
  value,
  unit = '',
  trend,
  height = 200,
  onDrillDown,
}: ChartWidgetProps) {
  const colors = [
    '#228be6',
    '#40c057',
    '#fab005',
    '#fa5252',
    '#be4bdb',
    '#15aabf',
    '#fd7e14',
    '#868e96',
  ];

  const renderBarChart = () => {
    if (!data || data.length === 0) return null;

    const maxValue = Math.max(...data.map((d) => d.value));

    return (
      <Stack gap="xs" h={height}>
        {data.map((item, index) => {
          const percentage = (item.value / maxValue) * 100;
          const color = item.color || colors[index % colors.length];

          return (
            <Group
              key={index}
              gap="xs"
              style={{ cursor: onDrillDown ? 'pointer' : 'default' }}
              onClick={() => onDrillDown?.(item.label)}
            >
              <Text size="xs" style={{ width: 80 }} truncate>
                {item.label}
              </Text>
              <Box
                style={{
                  flex: 1,
                  height: 20,
                  backgroundColor: '#f1f3f5',
                  borderRadius: 4,
                  position: 'relative',
                }}
              >
                <Box
                  style={{
                    width: `${percentage}%`,
                    height: '100%',
                    backgroundColor: color,
                    borderRadius: 4,
                    transition: 'width 0.3s ease',
                  }}
                />
              </Box>
              <Text size="xs" style={{ width: 50 }} ta="right">
                {item.value}
              </Text>
            </Group>
          );
        })}
      </Stack>
    );
  };

  const renderPieChart = () => {
    if (!data || data.length === 0) return null;

    const total = data.reduce((sum, d) => sum + d.value, 0);
    let currentAngle = 0;

    return (
      <Box h={height} style={{ position: 'relative' }}>
        <svg viewBox="0 0 100 100" style={{ width: '100%', height: '100%' }}>
          {data.map((item, index) => {
            const angle = (item.value / total) * 360;
            const startAngle = currentAngle;
            const endAngle = currentAngle + angle;
            currentAngle = endAngle;

            const startRad = (startAngle - 90) * (Math.PI / 180);
            const endRad = (endAngle - 90) * (Math.PI / 180);

            const x1 = 50 + 40 * Math.cos(startRad);
            const y1 = 50 + 40 * Math.sin(startRad);
            const x2 = 50 + 40 * Math.cos(endRad);
            const y2 = 50 + 40 * Math.sin(endRad);

            const largeArc = angle > 180 ? 1 : 0;
            const color = item.color || colors[index % colors.length];

            return (
              <path
                key={index}
                d={`M 50 50 L ${x1} ${y1} A 40 40 0 ${largeArc} 1 ${x2} ${y2} Z`}
                fill={color}
                style={{ cursor: 'pointer' }}
                onClick={() => onDrillDown?.(item.label)}
              />
            );
          })}
        </svg>
        <Group
          justify="center"
          style={{
            position: 'absolute',
            top: '50%',
            left: '50%',
            transform: 'translate(-50%, -50%)',
          }}
        >
          <Text size="lg" fw={500}>
            {total}
          </Text>
        </Group>
      </Box>
    );
  };

  const renderMetric = () => (
    <Stack align="center" justify="center" h={height}>
      <Text size="xl" fw={700}>
        {value !== undefined ? value.toLocaleString() : '-'}
        {unit && <Text component="span" size="md"> {unit}</Text>}
      </Text>
      {trend !== undefined && (
        <Text
          size="sm"
          c={trend >= 0 ? 'green' : 'red'}
        >
          {trend >= 0 ? '↑' : '↓'} {Math.abs(trend).toFixed(1)}%
        </Text>
      )}
    </Stack>
  );

  return (
    <Paper shadow="sm" p="md" withBorder h="100%">
      <Stack gap="md">
        <Text fw={500}>{title}</Text>
        {type === 'bar' && renderBarChart()}
        {type === 'pie' && renderPieChart()}
        {type === 'metric' && renderMetric()}
      </Stack>
    </Paper>
  );
}

import { useState, useEffect } from 'react';
import {
  AppShell,
  Navbar,
  Header,
  Footer,
  Text,
  Group,
  Button,
  ActionIcon,
  Tooltip,
  Burger,
  MediaQuery,
  useMantineTheme,
  Box,
  NavLink,
  Badge,
} from '@mantine/core';
import {
  IconDashboard,
  IconFileText,
  IconMapPin,
  IconChartBar,
  IconAlertTriangle,
  IconSettings,
  IconDatabase,
  IconCommand,
  IconRefresh,
  IconChevronRight,
} from '@tabler/icons-react';
import { useLocation, useNavigate } from 'react-router-dom';
import { CommandPalette } from '../components/CommandPalette';

interface AppLayoutProps {
  children: React.ReactNode;
}

export function AppLayout({ children }: AppLayoutProps) {
  const [opened, setOpened] = useState(false);
  const [commandPaletteOpen, setCommandPaletteOpen] = useState(false);
  const theme = useMantineTheme();
  const location = useLocation();
  const navigate = useNavigate();

  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
        e.preventDefault();
        setCommandPaletteOpen(true);
      }
    };

    window.addEventListener('keydown', handleKeyDown);
    return () => window.removeEventListener('keydown', handleKeyDown);
  }, []);

  const navItems = [
    {
      label: '统计驾驶舱',
      icon: <IconDashboard size={20} />,
      path: '/',
    },
    {
      label: '投诉单管理',
      icon: <IconFileText size={20} />,
      path: '/complaints',
    },
    {
      label: '采样点位工作台',
      icon: <IconMapPin size={20} />,
      path: '/sampling-points',
    },
    {
      label: '噪声读数',
      icon: <IconChartBar size={20} />,
      path: '/noise-readings',
    },
    {
      label: '异常分诊',
      icon: <IconAlertTriangle size={20} />,
      path: '/anomalies',
      badge: 3,
    },
    {
      label: '规则配置',
      icon: <IconSettings size={20} />,
      path: '/rules',
    },
    {
      label: 'Seed数据',
      icon: <IconDatabase size={20} />,
      path: '/seed',
    },
  ];

  const commands = [
    {
      id: 'dashboard',
      label: '打开统计驾驶舱',
      category: '导航',
      shortcut: 'G+D',
      action: () => navigate('/'),
    },
    {
      id: 'complaints',
      label: '打开投诉单管理',
      category: '导航',
      shortcut: 'G+C',
      action: () => navigate('/complaints'),
    },
    {
      id: 'sampling',
      label: '打开采样点位工作台',
      category: '导航',
      shortcut: 'G+S',
      action: () => navigate('/sampling-points'),
    },
    {
      id: 'readings',
      label: '打开噪声读数',
      category: '导航',
      shortcut: 'G+R',
      action: () => navigate('/noise-readings'),
    },
    {
      id: 'anomalies',
      label: '打开异常分诊',
      category: '导航',
      shortcut: 'G+A',
      action: () => navigate('/anomalies'),
    },
  ];

  return (
    <>
      <AppShell
        navbarOffsetBreakpoint="sm"
        asideOffsetBreakpoint="sm"
        navbar={
          <Navbar
            p="md"
            hiddenBreakpoint="sm"
            hidden={!opened}
            width={{ base: 250 }}
          >
            <Navbar.Section>
              <Stack gap="xs">
                {navItems.map((item) => (
                  <NavLink
                    key={item.path}
                    label={item.label}
                    icon={item.icon}
                    rightSection={
                      item.badge ? (
                        <Badge size="xs" variant="filled">
                          {item.badge}
                        </Badge>
                      ) : location.pathname === item.path ? (
                        <IconChevronRight size={12} />
                      ) : null
                    }
                    active={location.pathname === item.path}
                    onClick={() => {
                      navigate(item.path);
                      setOpened(false);
                    }}
                    variant="filled"
                  />
                ))}
              </Stack>
            </Navbar.Section>

            <Navbar.Section mt="auto">
              <Button
                fullWidth
                variant="light"
                leftSection={<IconCommand size={16} />}
                onClick={() => setCommandPaletteOpen(true)}
              >
                命令面板
              </Button>
            </Navbar.Section>
          </Navbar>
        }
        header={
          <Header height={{ base: 50, md: 70 }} p="md">
            <Group sx={{ height: '100%' }} px="md" position="apart">
              <MediaQuery largerThan="sm" styles={{ display: 'none' }}>
                <Burger
                  opened={opened}
                  onClick={() => setOpened((o) => !o)}
                  size="sm"
                  color={theme.colors.gray[6]}
                />
              </MediaQuery>

              <Text size="lg" fw={700}>
                环境噪声投诉采样证据归档控制台
              </Text>

              <Group>
                <Tooltip label="命令面板 (Ctrl+K)">
                  <ActionIcon
                    variant="subtle"
                    onClick={() => setCommandPaletteOpen(true)}
                  >
                    <IconCommand size={18} />
                  </ActionIcon>
                </Tooltip>
                <Tooltip label="刷新数据">
                  <ActionIcon variant="subtle">
                    <IconRefresh size={18} />
                  </ActionIcon>
                </Tooltip>
              </Group>
            </Group>
          </Header>
        }
      >
        <Box p="md">{children}</Box>
      </AppShell>

      <CommandPalette
        opened={commandPaletteOpen}
        onClose={() => setCommandPaletteOpen(false)}
        commands={commands}
      />
    </>
  );
}

import { Stack } from '@mantine/core';

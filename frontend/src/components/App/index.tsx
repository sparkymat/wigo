/* eslint-disable react/no-unknown-property */
import React, { useCallback, useEffect, useState } from 'react';
import { Routes, Route, useLocation, useNavigate } from 'react-router-dom';
import {
  ActionIcon,
  Anchor,
  AppShell,
  Burger,
  ColorScheme,
  Flex,
  Header,
  MediaQuery,
  Menu,
  NavLink,
  Navbar,
  Space,
  Title,
  useMantineTheme,
} from '@mantine/core';
import { useDispatch, useSelector } from 'react-redux';
import { useLocalStorage } from '@mantine/hooks';
import {
  IconBrightness,
  IconDashboard,
  IconUserCircle,
  IconUsers,
} from '@tabler/icons-react';

import { AppDispatch } from '../../store';
import { updatePath } from '../../features/App/slice';
import { selectPath } from '../../selectors/App';
import Home from '../Home';

interface Path {
  href: string;
  label: string;
  icon: React.JSX.Element;
}

const App = () => {
  const dispatch = useDispatch<AppDispatch>();
  const [opened, setOpened] = useState(false);
  const theme = useMantineTheme();

  const location = useLocation();
  const navigate = useNavigate();

  const currentPath = useSelector(selectPath);

  const [colorScheme, setColorScheme] = useLocalStorage<ColorScheme>({
    key: 'wigo-color-scheme',
    defaultValue: 'light',
    getInitialValueInEffect: true,
  });

  const toggleColorScheme = (value?: ColorScheme) =>
    setColorScheme(value || (colorScheme === 'dark' ? 'light' : 'dark'));

  const paths: Path[] = [
    {
      href: '/',
      label: 'Dashboard',
      icon: <IconDashboard size="1rem" stroke={1.5} />,
    },
    {
      href: '/feeds',
      label: 'Feeds',
      icon: <IconUsers size="1rem" stroke={1.5} />,
    },
  ];

  const onNavClick = useCallback(
    (p: string) => {
      dispatch(updatePath(p));
      setOpened(false);
      window.location.href = `/#${p}`;
    },
    [dispatch],
  );

  const isActive = useCallback(
    (path: string): boolean => currentPath === path,
    [currentPath],
  );

  useEffect(() => {
    dispatch(updatePath(location.pathname));
  }, [dispatch, location]);

  return (
    <div>
      <AppShell
        padding="md"
        navbar={
          <Navbar
            p="md"
            hiddenBreakpoint="sm"
            hidden={!opened}
            width={{ sm: 200, lg: 300 }}
          >
            {paths.map(p => (
              <NavLink
                label={p.label}
                active={isActive(p.href)}
                variant="subtle"
                tt="uppercase"
                icon={p.icon}
                onClick={() => onNavClick(p.href)}
              />
            ))}
          </Navbar>
        }
        header={
          <Header height={{ base: 50, md: 70 }} p="md">
            <div
              style={{ display: 'flex', alignItems: 'center', height: '100%' }}
            >
              <MediaQuery largerThan="sm" styles={{ display: 'none' }}>
                <Burger
                  opened={opened}
                  onClick={() => setOpened(o => !o)}
                  size="sm"
                  color={theme.colors.gray[6]}
                  mr="xl"
                />
              </MediaQuery>
              <Flex
                align="center"
                direction="row"
                wrap="wrap"
                sx={{ width: '100%' }}
              >
                <Anchor
                  underline={false}
                  component="button"
                  onClick={() => navigate('/')}
                >
                  <Title order={3} weight={300}>
                    what&apos;s going on
                  </Title>
                </Anchor>

                <Space sx={{ flex: 1 }} w="sm" />

                <Menu shadow="md" width={200}>
                  <Menu.Target>
                    <ActionIcon>
                      <IconUserCircle size={24} />
                    </ActionIcon>
                  </Menu.Target>

                  <Menu.Dropdown>
                    <Menu.Item
                      icon={<IconBrightness size={14} />}
                      onClick={() => toggleColorScheme()}
                    >
                      Toggle dark mode
                    </Menu.Item>
                  </Menu.Dropdown>
                </Menu>
              </Flex>
            </div>
          </Header>
        }
        styles={thisTheme => ({
          main: {
            backgroundColor:
              thisTheme.colorScheme === 'dark'
                ? thisTheme.colors.dark[8]
                : thisTheme.colors.gray[0],
          },
        })}
      >
        <Routes>
          <Route index element={<Home />} />
        </Routes>
      </AppShell>
    </div>
  );
};

export default App;

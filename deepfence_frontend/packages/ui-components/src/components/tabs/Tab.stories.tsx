import { Meta, StoryObj } from '@storybook/react';
import { useState } from 'react';
import { FaAdn, FaAffiliatetheme, FaAirbnb } from 'react-icons/fa';

import Tab from '@/components/tabs/Tabs';

export default {
  title: 'Components/Tab',
  component: Tab,
  argTypes: {
    onValueChange: { action: 'onValueChange' },
  },
} satisfies Meta<typeof Tab>;

const tabs = [
  {
    label: 'Tab One',
    value: 'Tab1',
  },
  {
    label: 'Tab Two',
    value: 'Tab2',
  },
  {
    label: 'Tab Three',
    value: 'Tab3',
  },
  {
    label: 'Some tab name',
    value: 'Tab4',
    disabled: true,
  },
];

export const Default: StoryObj<typeof Tab> = {
  args: {
    tabs,
  },
};

const tabs2 = [
  {
    label: 'Tab One',
    value: 'tab1',
    icon: <FaAdn />,
  },
  {
    label: 'Tab Two',
    value: 'tab2',
    icon: <FaAffiliatetheme />,
  },
  {
    label: 'Tab Three',
    value: 'tab3',
    icon: <FaAirbnb />,
  },
];
const WithContent = () => {
  const [tab, setTab] = useState('tab1');
  return (
    <Tab value={tab} defaultValue={tab} tabs={tabs2} onValueChange={(v) => setTab(v)}>
      <div className="h-full p-2 dark:text-white">
        You are now on {tabs2.find((t) => t.value === tab)?.label}
      </div>
    </Tab>
  );
};

export const TabWithContent: StoryObj<typeof Tab> = {
  render: WithContent,
};

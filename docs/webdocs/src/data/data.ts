export type tabItem = {
  name: string;
  path: string;
};

export type tab = {
  name: string;
  collapsed: boolean;
  tabItems: tabItem[];
};

export const tabs: tab[] = [
  {
    name: "Start Here",
    collapsed: false,
    tabItems: [
      {
        name: "Getting Started",
        path: "/[lang]/[version]/getting_started",
      },
      {
        name: "Why Make This?",
        path: "/[lang]/[version]/why",
      },
    ],
  },
  {
    name: "Routes",
    collapsed: false,
    tabItems: [
      {
        name: "Quarters",
        path: "/[lang]/[version]/quarters",
      },
      {
        name: "Subjects",
        path: "/[lang]/[version]/subjects",
      },
      {
        name: "Courses",
        path: "/[lang]/[version]/courses",
      },
      {
        name: "Sections",
        path: "/[lang]/[version]/sections",
      },
    ],
  }
];

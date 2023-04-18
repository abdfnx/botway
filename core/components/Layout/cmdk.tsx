import { useState } from "react";
import CommandPalette, {
  JsonStructure,
  filterItems,
  renderJsonStructure,
  useHandleOpenCommandPalette,
} from "react-cmdk";

export const CMDK = ({ isOpen, setIsOpen }: any) => {
  const [selected, setSelected] = useState<number>(0);
  const [search, setSearch] = useState<string>("");
  const [page, setPage] = useState<"root" | "positions">("root");

  useHandleOpenCommandPalette(setIsOpen);

  const items: JsonStructure = [
    {
      heading: "Home",
      id: "home",
      items: [
        {
          children: "Home",
          icon: "HomeIcon",
          id: "home",
          disabled: true,
          href: "#",
          renderLink: (props) => <a {...props} />,
        },
        {
          children: "Settings",
          icon: "CogIcon",
          id: "settings",
          disabled: true,
        },
        {
          children: "Positions",
          icon: "BriefcaseIcon",
          closeOnSelect: false,
          keywords: ["jobs"],
          id: "positions",
          onClick: () => {
            setPage("positions");
            setSearch("");
          },
        },
        {
          children: "Candidates",
          icon: "UsersIcon",
          id: "users",
          onClick: () => {
            alert("hj");
          },
        },
      ],
    },
    {
      heading: "External",
      id: "external",
      items: [
        {
          href: "https://google.com",
          children: "Help",
          icon: "LifebuoyIcon",
          id: "support",
          target: "_blank",
          rel: "noopener noreferrer",
        },
      ],
    },
    {
      heading: "Extra",
      id: "extra",
      items: [
        {
          children: "Privacy policy",
          icon: "FlagIcon",
          id: "privacy",
        },
        {
          children: "User agreement",
          icon: "UserIcon",
          id: "user-agreement",
        },
        {
          children: "About",
          icon: "EyeIcon",
          id: "about",
        },
        {
          children: "Career",
          icon: "UsersIcon",
          id: "career",
        },
      ],
    },
  ];

  const rootItems = filterItems(items, search);

  return (
    <CommandPalette
      onChangeSelected={setSelected}
      onChangeSearch={setSearch}
      onChangeOpen={setIsOpen}
      selected={selected}
      search={search}
      isOpen={isOpen}
      page={page}
    >
      <CommandPalette.Page id="root" searchPrefix={["General"]}>
        {rootItems.length ? renderJsonStructure(rootItems) : <p>Not Found</p>}
      </CommandPalette.Page>

      <CommandPalette.Page
        searchPrefix={["General", "Positions"]}
        id="positions"
        onEscape={() => {
          setPage("root");
        }}
      >
        <CommandPalette.List heading="Positions">
          <CommandPalette.ListItem index={0}>
            Nothing here
          </CommandPalette.ListItem>
        </CommandPalette.List>
      </CommandPalette.Page>
    </CommandPalette>
  );
};

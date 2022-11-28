import Avatar from "boring-avatars";

export const UserAvatar = ({ data, size }: any) => {
  return <Avatar size={size} name={data} variant="marble" />;
};

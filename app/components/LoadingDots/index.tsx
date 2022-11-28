import clsx from "clsx";
import styles from "./ld.module.scss";

export const LoadingDots = ({ children, className }: any) => {
  return (
    <span className={clsx(styles.loading, className)}>
      {children && <div className={styles.child}>{children}</div>}
      <span />
      <span />
      <span />
    </span>
  );
};

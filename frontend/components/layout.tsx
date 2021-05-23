import Head from 'next/head';
import styles from '../styles/Home.module.css';

type PageMeta = {
  title: string;
  description: string;
};
const Layout: React.FC<PageMeta> = ({ title, description, children }) => {
  return (
    <div className={styles.container}>
      <Head>
        <title>{title}</title>
        <meta name="description" content={description} />
      </Head>
      {children}
    </div>
  );
};
export default Layout;

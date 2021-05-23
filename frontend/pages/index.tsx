import Layout from '../components/layout';
import styles from '../styles/Home.module.css';

const Home: React.FC = () => {
  return (
    <Layout title="Image Uploader" description="画像の共有ができます">
      <main className={styles.main}>ここ</main>
      <footer className={styles.footer}>ここ</footer>
    </Layout>
  );
};
export default Home;

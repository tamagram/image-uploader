import Layout from '../components/layout';
import styles from '../styles/Home.module.css';
import Image from 'next/image'
import { ChangeEvent, useState } from 'react';
import axios from 'axios';

const Home: React.FC = () => {
  const [image, setImage] = useState('');
  const [message, setMessage] = useState('画像を選択してください')

  const fileSelect = (event: ChangeEvent<any>) => {
    const imageFile = event.target.files[0];
    const imageUrl = URL.createObjectURL(imageFile);
    setImage(imageUrl);
    setMessage('');
  };
  const fileSend = () => {
    setMessage('送信中...');
    const params = new FormData();
    params.append('imageFile', image);
    axios
      .post(
        process.env.BACKEND_URL as string,
        params,
        {
          headers: {
            'content-type': 'multipart/data'
          }
        },
      )
      .then((result) => {
        setMessage('送信完了');
        console.log(result);
      })
      .catch((reason) => {
        setMessage('送信失敗');
        console.log(reason);
      });
  };

  return (
    <Layout title="Image Uploader" description="画像の共有ができます">
      <main className={styles.main}>
        <h1>Image Uploader</h1>
        <form method='post' encType='multipart/form-data'>
          <input type='file' accept='image/*' onChange={fileSelect}></input>
          <input type='button' value='send' onClick={fileSend}></input>
        </form>
        {message}
        <img className={styles.image} src={image}></img>
      </main>
      <footer className={styles.footer}>
        <a
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by{' '}
          <span className={styles.logo}>
            <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16} />
          </span>
        </a>
      </footer>
    </Layout>
  );
};
export default Home;

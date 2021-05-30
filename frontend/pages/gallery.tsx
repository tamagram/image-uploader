import { GetServerSideProps } from "next";
import { imageFetch } from "../lib/api";

type FetchImage = {
  result: any
}

const Gallery: React.FC<FetchImage> = ({ result }) => {
  
  console.log(new FormData(result))
  return (
    <div>
      <h1>Gallery</h1>
      {/* <img src={galleryImage}></img> */}
    </div>
  )
};
export default Gallery

// imagefetch
export const getServerSideProps: GetServerSideProps = async (context) => {
  const result = await imageFetch()
  return {
    props: { result },
  }
};

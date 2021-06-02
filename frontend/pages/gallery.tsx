import axios from "axios";
import React, { useState } from "react";
import useSWR from "swr";
import Layout from "../components/layout";

type Image = {
  name: string
  data: string
}

const axiosFetcher = async () => {
  const result = await axios.get<Image[]>(
    'http://localhost:5000/images'
  )
  return result.data
}

const Gallery: React.FC = () => {
  const { data: images, error } = useSWR('imagesFetch', axiosFetcher)
  if (error) return <div>failed to load</div>
  if (!images) {
    return <div>loading...</div>
  }
  return (
    <Layout title="Gallery" description="画像の一覧">
      <div>
        <h2>Gallery</h2>
        <ul>
          {images &&
            images.map((image, index) => {
              return (
                <li key={index}>
                  <p>{image.name}</p>
                  <img src={"data:image/*;base64," + image.data}></img>
                </li>
              )
            })
          }
        </ul>
      </div>
    </Layout>
  )
};
export default Gallery

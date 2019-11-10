import React from "react";
import HeroSection from "./../components/HeroSection";
import FeaturesSection from "./../components/FeaturesSection";
import { useRouter } from "./../util/router.js";

function IndexPage(props) {
  const router = useRouter();

  return (
    <>
      <HeroSection
        color="white"
        size="medium"
        title="GOLANG JOBS"
        subtitle="ПУБЛИКАЦИЯ БЕСПЛАТНА И НЕ ТРЕБУЕТ РЕГИСТРАЦИИ"
        buttonText="ДОБАВИТЬ ВАКАНСИЮ"
        image="https://www.mtoag.com/assets/images/Golang-Go.svg"
        buttonOnClick={() => {
          router.push("/jobs/new");
        }}
      />
      <hr/>
      <FeaturesSection
        color="white"
        size="medium"
        title="АКТИВНЫЕ ВАКАНСИИ"
        subtitle=""
      />
    </>
  );
}

export default IndexPage;

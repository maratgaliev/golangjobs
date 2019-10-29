import React from "react";
import ContentSection from "./../components/ContentSection";
import AboutSection from "./../components/AboutSection";

function AboutPage(props) {
  return (
    <>
      <ContentSection
        color="info"
        size="large"
        title="О ПРОЕКТЕ"
        subtitle="ИНФОРМАЦИЯ"
      />
      <AboutSection
        color="white"
        size="medium"
        title=""
        subtitle=""
      />
    </>
  );
}

export default AboutPage;

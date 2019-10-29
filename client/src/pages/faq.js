import React from "react";
import FaqSection from "./../components/FaqSection";
import ContentSection from "./../components/ContentSection";

function FaqPage(props) {
  return (
    <>
      <ContentSection
      color="info"
      size="large"
      title="FAQ"
      subtitle="ЧАСТО ЗАДАВАЕМЫЕ ВОПРОСЫ"
    />
    <FaqSection
      color="white"
      size="medium"
      title=""
      subtitle=""
    />
    </>
  );
}

export default FaqPage;

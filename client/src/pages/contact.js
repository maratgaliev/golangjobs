import React from "react";
import ContentSection from "./../components/ContentSection";
import ContactSection from "./../components/ContactSection";

function ContactPage(props) {
  return (
    <>
      <ContentSection
        color="info"
        size="large"
        title="КОНТАКТЫ"
        subtitle="КОНТАКТНЫЕ ДАННЫЕ"
      />
      <ContactSection
        color="white"
        size="medium"
        title=""
        subtitle=""
      />
    </>
  );
}

export default ContactPage;

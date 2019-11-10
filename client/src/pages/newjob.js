import React from "react";
import NewJobSection from "./../components/NewJobSection";

function NewJobPage(props) {
  return (
    <NewJobSection
      color="white"
      size="medium"
      title="Новая вакансия"
      subtitle=""
      buttonText="Отправить на модерацию"
    />
  );
}

export default NewJobPage;

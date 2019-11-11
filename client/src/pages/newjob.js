import React from "react";
import NewJobSection from "./../components/NewJobSection";
import { withRouter } from 'react-router-dom';

function NewJobPage(props) {
  return (
    <NewJobSection
      color="white"
      size="medium"
      title="Новая вакансия"
      subtitle=""
      message="Спасибо, вакансия успешно добавлена и появится на сайте после модерации"
      buttonText="Отправить на модерацию"
    />
  );
}

export default withRouter(NewJobPage);

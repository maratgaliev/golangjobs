import React from "react";
import Section from "./Section";

function AboutSection(props) {
  return (
    <Section color={props.color} size={props.size}>
      <div className="container">
        <div className="content is-medium has-text-left">
          <p>Golangjobs.ru является некоммерческим open-source проектом.</p> 
          <p>Основная цель и идея разработки - популяризация технологии и поддержка сообщества разработчиков.</p>
          <p>Все вакансии публикуются бесплатно. Размещаемая информация передается "as is" и принадлежит ее авторам.</p>
          <p><a href="mailto:kazanlug@gmail.com">Свяжитесь</a> со мной, если у вас есть идеи, предложения и вопросы.</p>
          <p>Код проекта на <a href='https://github.com/maratgaliev/golangjobs'>GitHub</a>.</p>
        </div>
      </div>
    </Section>
  );
}

export default AboutSection;

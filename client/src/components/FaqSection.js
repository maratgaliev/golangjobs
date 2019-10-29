import React from "react";
import Section from "./Section";
import SectionHeader from "./SectionHeader";
import Faq from "./Faq";

function FaqSection(props) {
  return (
    <Section color={props.color} size={props.size}>
      <div className="container">
        <SectionHeader
          title={props.title}
          subtitle={props.subtitle}
          centered={true}
          size={3}
        />
        <Faq
          items={[
            {
              question: "Для кого этот сайт?",
              answer:
                "Сайт призван облегчить поиск работы связанной с языком программирования Golang и создан, как для работодателей, так и для соискателей."
            },
            {
              question: "Где взять исходный код приложения?",
              answer:
                "На github - https://github.com/maratgaliev/golangjobs"
            },
            {
              question: "Могу ли я переиспользовать код приложения?",
              answer:
                "Да, можно переиспользовать код в любых целях"
            },
            {
              question: "Как связаться с автором?",
              answer:
                "На сайте есть специальный раздел с контактами, вы можете написать на почту kazanlug@gmail.com"
            }
          ]}
        />
      </div>
    </Section>
  );
}

export default FaqSection;

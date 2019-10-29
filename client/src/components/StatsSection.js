import React from "react";
import Section from "./Section";
import "./StatsSection.scss";

function StatsSection(props) {
  let items = [
    {
    	title: "ВАКАНСИЙ",
    	stat: "101"
    },
    {
    	title: "КОМПАНИЙ",
    	stat: "102"
    },
    {
    	title: "СОБЫТИЙ",
    	stat: "103"
    },
    {
    	title: "ОТЗЫВОВ",
    	stat: "104"
    }
  ]
  return (
    <Section color={props.color} size={props.size}>
      <nav className="StatsSection__level level">
        {items.map((item, index) => (
          <div className="level-item has-text-centered" key={index}>
            <div>
              <p className="heading">{item.title}</p>
              <p className="title">{item.stat}</p>
            </div>
          </div>
        ))}
      </nav>
    </Section>
  );
}

export default StatsSection;

import React from "react";
import HeroSection from "./../components/HeroSection";
import ClientsSection from "./../components/ClientsSection";
import FeaturesSection from "./../components/FeaturesSection";
import NewsletterSection from "./../components/NewsletterSection";
import { useRouter } from "./../util/router.js";
import StatsSection from "../components/StatsSection";

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
          router.push("/");
        }}
      />
      <ClientsSection color="light" size="normal" title="КОМПАНИИ" subtitle="" />
      <StatsSection color="white" size= "medium" />
      <FeaturesSection
        color="white"
        size="medium"
        title="АКТИВНЫЕ ВАКАНСИИ"
        subtitle=""
      />
      <NewsletterSection
        color="white"
        size="medium"
        title="Подписка на новости"
        subtitle="Обновления сервиса в ваш почтовый ящик"
        buttonText="Подписаться"
        inputPlaceholder="Введите email"
        subscribedMessage="Вы успешно подписаны на рассылку"
      />
    </>
  );
}

export default IndexPage;

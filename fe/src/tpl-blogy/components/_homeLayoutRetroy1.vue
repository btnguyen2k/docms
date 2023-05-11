<template>
  <!-- Start retroy layout blog posts/style 1 -->
  <section class="section bg-light">
    <div class="container">
      <div class="row align-items-stretch retro-layout">
        <div class="col-md-4">
          <a href="single.html" class="h-entry mb-30 v-height gradient">
            <div class="featured-img" id="bgimg1"></div>
            <!--<div class="featured-img" style="background-image: url('images/img_2_horizontal.jpg');"></div>-->
<!--            <div class="featured-img" style="background-image: url('//placehold.co/964x964/31343C/90A1A9?font=lato&text=Lato');"></div>-->
            <div class="text">
              <span class="date">Apr. 14th, 2022</span>
              <h2>AI can now kill those annoying cookie pop-ups</h2>
            </div>
          </a>
          <a href="single.html" class="h-entry v-height gradient">
            <div class="featured-img" id="bgimg2"></div>
            <!--<div class="featured-img" style="background-image: url('images/img_5_horizontal.jpg');"></div>-->
<!--            <div class="featured-img" style="background-image: url('//placehold.co/964x964/31343C/90A1A9?font=lora&text=Lora');"></div>-->
            <div class="text">
              <span class="date">Apr. 14th, 2022</span>
              <h2>Don’t assume your user data in the cloud is safe</h2>
            </div>
          </a>
        </div>
        <div class="col-md-4">
          <a href="single.html" class="h-entry img-5 h-100 gradient">
            <div class="featured-img" id="bgimg0"></div>
            <!--<div class="featured-img" style="background-image: url('images/img_1_vertical.jpg');"></div>-->
<!--            <div class="featured-img" style="background-image: url('//placehold.co/964x964/31343C/90A1A9?font=montserrat&text=Montserrat');"></div>-->
            <div class="text">
              <span class="date">Apr. 14th, 2022</span>
              <h2>Why is my internet so slow?</h2>
            </div>
          </a>
        </div>
        <div class="col-md-4">
          <a href="single.html" class="h-entry mb-30 v-height gradient">
            <div class="featured-img" id="bgimg3"></div>
            <!--<div class="featured-img" style="background-image: url('images/img_3_horizontal.jpg');"></div>-->
<!--            <div class="featured-img" style="background-image: url('//placehold.co/964x964/31343C/90A1A9?font=open-sans&text=Open+Sans');"></div>-->
            <div class="text">
              <span class="date">Apr. 14th, 2022</span>
              <h2>Startup vs corporate: What job suits you best?</h2>
            </div>
          </a>
          <a href="single.html" class="h-entry v-height gradient">
            <div class="featured-img" id="bgimg4"></div>
            <!--<div class="featured-img" style="background-image: url('images/img_4_horizontal.jpg');"></div>-->
<!--            <div class="featured-img" style="background-image: url('//placehold.co/964x964/31343C/90A1A9?font=oswald&text=Oswald');"></div>-->
            <div class="text">
              <span class="date">Apr. 14th, 2022</span>
              <h2>Thought you loved Python? Wait until you meet Rust</h2>
            </div>
          </a>
        </div>
      </div>
    </div>
  </section>
  <!-- End retroy layout blog posts -->
</template>

<style lang="css">
#bgimg0 {
  background-image: url('@/tpl-blogy/assets/images/img_1_vertical.jpg');
}
#bgimg1 {
  background-image: url('@/tpl-blogy/assets/images/img_2_horizontal.jpg');
}
#bgimg2 {
  background-image: url('@/tpl-blogy/assets/images/img_5_horizontal.jpg');
}
#bgimg3 {
  background-image: url('@/tpl-blogy/assets/images/img_3_horizontal.jpg');
}
#bgimg4 {
  background-image: url('@/tpl-blogy/assets/images/img_4_horizontal.jpg');
}
</style>

<script>
export default {
  name: 'lego-home-layout-retroy1',
  inject: ['$siteMeta'],
}
</script>

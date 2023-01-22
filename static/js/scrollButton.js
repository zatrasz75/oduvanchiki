const goTopBtn = document.querySelector(".button");

goTopBtn.addEventListener("click", goTop);

function goTop() {

  if (window.pageYOffset > 0) {
    window.scrollBy(0, -75); // второй аргумент - скорость
    setTimeout(goTop, 0); // входим в рекурсию
  }
}
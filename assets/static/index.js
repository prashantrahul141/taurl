const formElement = document.getElementById("FORM");
const inputElement = document.getElementById("INPUT");
const shortendElement = document.getElementById("SHORTEND");
const shortendTextElement = document.getElementById("SHORTEND_TEXT");
const buttonElement = document.getElementById("SUBMIT_BUTTON");

shortendElement.style.opacity = 0;

formElement.addEventListener("submit", async (e) => {
  inputElement.disabled = true;
  buttonElement.disabled = true;

  e.preventDefault();
  const original_url = e.target.elements.original_url.value;
  await onShortendUrl(original_url);

  inputElement.disabled = false;
  buttonElement.disabled = false;
});

const onShortendUrl = async (original_url) => {
  shortendElement.style.opacity = 100;
  original_url = original_url.trim();

  if (original_url.length <= 0) {
    console.log("ERROR::VALIDATION FAILED.");
  }

  const res = await fetch(window.location.origin + "/api/set", {
    method: "POST",
    body: JSON.stringify({ Url: original_url }),
  });

  if (res.status != 201) {
    console.error(res);
    return;
  }

  const jsonData = await res.json();
  shortendTextElement.innerText = jsonData.ShortendUrl;
};

const copyUrl = (e) => {
  if (!navigator.clipboard) {
    alert("Your browser doesn't support copying to clipboard.");
    return;
  }

  navigator.clipboard.writeText(e.target.innerText).then(
    function () {},
    function (err) {
      console.error("failed copying.", err);
    },
  );
};

shortendElement.addEventListener("click", copyUrl);

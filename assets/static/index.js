const formElement = document.getElementById("FORM");
formElement.addEventListener("submit", async (e) => {
  e.preventDefault();
  const original_url = e.target.elements.original_url.value;
  await onShortendUrl(original_url);
});

const onShortendUrl = async (original_url) => {
  original_url = original_url.trim();

  if (original_url.length <= 0) {
    console.log("ERROR::VALIDATION FAILED.");
  }

  const res = await fetch(window.location.origin + "/api/set", {
    method: "POST",
    body: JSON.stringify({ Url: original_url }),
  });

  console.log(res);
  const jsonData = await res.json();
  console.log(jsonData);
};

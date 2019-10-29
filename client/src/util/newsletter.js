function subscribe(data) {
  return fetch("/.netlify/functions/newsletter", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  }).then(r => r.json());
}

export default { subscribe };

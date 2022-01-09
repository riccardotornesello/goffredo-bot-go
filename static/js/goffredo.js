function deleteSound(id) {
  const formData = new FormData();
  formData.append("id", id);

  fetch(window.location.href, {
    method: "delete",
    body: formData,
  }).then(() => {
    location.reload();
  });
}

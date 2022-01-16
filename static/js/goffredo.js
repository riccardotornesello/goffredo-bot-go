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

var deleteModal = document.getElementById("deleteModal");
deleteModal.addEventListener("show.bs.modal", function (event) {
  var button = event.relatedTarget;
  var id = button.getAttribute("data-bs-id");
  var name = button.getAttribute("data-bs-name");

  var modalBody = deleteModal.querySelector(".modal-body");
  var submitButton = deleteModal.querySelector("#submit-button");

  modalBody.textContent = `Do you want to delete the sound ${name}?`;
  submitButton.onclick = function () {
    submitButton.firstElementChild.classList.remove("visually-hidden");
    deleteSound(id);
  };
});

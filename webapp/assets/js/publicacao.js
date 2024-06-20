$("#nova-publicacao").on("submit", publicacao);

function publicacao(e) {
  e.preventDefault();

  if ($("#titulo").val() === "") {
    alert("Digite titulo...");
    $("#titulo").text("").focus();
    return;
  }
  if ($("#conteudo").val() === "") {
    alert("Digie conteudo...");
    $("#conteudo").text("").focus();
    return;
  }

  $.ajax({
    url: "/publicacoes",
    method: "POST",
    data: {
      titulo: $("#titulo").val(),
      conteudo: $("#conteudo").val(),
    },
  })
    .done(() => {
      window.location.reload();
    })
    .fail((erro) => {
      console.log(erro);
      alert("Erro na publicação");
    });
}

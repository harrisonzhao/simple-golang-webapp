$(function() {
  var todoCards = [];

  var getCardHtml = function(todo) {
    return `
  <div class="card">
    <div class="card-block">
      <button type="button" class="close js-close"><span>&times;</span></button>
      <div class="id display-none">` + todo.id + `</div>
      <h4 contenteditable class="card-title">` + todo.name + `</h4>
    </div>
  </div>`;
  };
  var getCardDeckHtml = function(cardsHtml) {
    return '<div class="card-deck">'+cardsHtml+'</div>';
  }
  var cardsPerRow = 3;
  var $cardDeckWrapper = $('.js-card-deck-wrapper');
  var renderCards = function() {
    $cardDeckWrapper.empty();
    var rowCardsHtml = '';
    var i = 0;
    for (i = 0; i < todoCards.length; ++i) {
      rowCardsHtml += getCardHtml(todoCards[i]);
      if (i > 0 && (i+1) % cardsPerRow === 0) {
        $cardDeckWrapper.append(getCardDeckHtml(rowCardsHtml));
        rowCardsHtml = '';
      }
    }
    if (i % cardsPerRow !== 0) {
      $cardDeckWrapper.append(getCardDeckHtml(rowCardsHtml));
    }
  };

  var createCard = function(todo) {
    todoCards.unshift(todo);
    renderCards();
  };
  var $todoAlert = $('.js-todo-input-alert');
  var $todoInput = $('#js-todo-input');
  var $todoAdd = $('.js-create-todo');
  // create card
  $todoAdd.click(function() {
    var name = $todoInput.val();
    if (name.length > 0) {
      var data = {
        name: name
      };
      $.post('/todos', JSON.stringify(data), function(todo) {
        $todoInput.val('');
        createCard(todo);
      });
    } else {
      $todoAlert.show();
      $todoAlert.fadeOut(3000, function() {});
    }
  });
  var enterKeyCode = 13;
  $todoInput.keydown(function(e) {
    if (e.keyCode === enterKeyCode) {
      $todoAdd.click();
    }
  });

  var listCards = function() {
    $.get('/todos', function(todos) {
      todoCards = todos.data;
      // show most recent cards first
      todoCards.reverse();
      renderCards();
    });
  };

  var updateCard = function(todo) {
    $.ajax({
      url: '/todos',
      type: 'PUT',
      data: JSON.stringify(todo),
      success: function() {
        $todoSuccess.show();
        $todoSuccess.fadeOut(3000, function() {});
      }
    });
  };
  var $todoSuccess = $('.js-todo-update-alert');
  // update card
  $cardDeckWrapper.on('keydown', 'h4', function(e) {
    if (e.keyCode === enterKeyCode) {
      e.preventDefault();
      var $this = $(this);
      var todo = {
        id: parseInt($this.parent().find('.id').text()),
        name: $this.text()
      };
      updateCard(todo);
    }
  });

  var deleteCard = function(todoId) {
    $.ajax({
      url: '/todos/' + todoId,
      type: 'DELETE',
      success: function() {
        todoCards = todoCards.filter(function(todo) {
          return todo.id !== todoId;
        });
        renderCards();
      }
    });
  };
  $cardDeckWrapper.on('click', '.js-close', function(e) {
    var id = parseInt($(this).parent().find('.id').text());
    deleteCard(id);
  });

  listCards();
});
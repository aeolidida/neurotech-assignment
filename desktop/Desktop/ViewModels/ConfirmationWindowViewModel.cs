using ReactiveUI;
using System;
using System.Reactive;
using Avalonia.Controls;

namespace Desktop.ViewModels
{
    public class ConfirmationWindowViewModel : ReactiveObject
    {
        private readonly Window _window;

        public ConfirmationWindowViewModel(Window window)
        {
            _window = window;

            ConfirmCommand = ReactiveCommand.Create(() =>
            {
                IsConfirmed = true;
                Close();
            });

            CancelCommand = ReactiveCommand.Create(() =>
            {
                IsConfirmed = false;
                Close();
            });
        }

        private bool _isConfirmed;
        public bool IsConfirmed
        {
            get => _isConfirmed;
            set => this.RaiseAndSetIfChanged(ref _isConfirmed, value);
        }

        public ReactiveCommand<Unit, Unit> ConfirmCommand { get; }
        public ReactiveCommand<Unit, Unit> CancelCommand { get; }

        private void Close()
        {
            _window.Close();
        }
    }
}
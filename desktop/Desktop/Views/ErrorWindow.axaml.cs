using System;
using System.Threading.Tasks;
using Avalonia.Controls;
using Avalonia.Markup.Xaml;
using Desktop.Models;
using Desktop.ViewModels;

namespace Desktop.Views
{
    public partial class ErrorWindow : Window
    {
        public ErrorWindow() 
        {
            InitializeComponent();
        }

        public ErrorWindow(string errorMessage)
        {
            InitializeComponent();
            var errorMessageControl = this.FindControl<TextBlock>("ErrorMessage");
            if (errorMessageControl != null)
            {
                errorMessageControl.Text = errorMessage;
            }
        }
        private void InitializeComponent()
        {
            AvaloniaXamlLoader.Load(this);
        }

        private void OkButton_Click(object? sender, Avalonia.Interactivity.RoutedEventArgs e)
        {
            Close();
        }
    }
}
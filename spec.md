# Specification for RFP Eval HTML App

This document outlines the detailed specification for the "RFP Eval" application, a single-page web application designed to assist in evaluating Request for Proposals (RFPs).

## 1. Project Overview

RFP Eval is a client-side chat application that allows users to interact with a large language model (LLM) to evaluate RFPs. The user can configure the application to use various OpenAI-compatible APIs (like OpenAI, Gemini, or a local Ollama instance).

The core workflow involves:
1.  Configuring the API provider, model, and a "System Prompt".
2.  Uploading an RFP document (in PDF format).
3.  Asking the AI to evaluate the RFP based on the provided System Prompt.
4.  The AI grades the fitment as High, Medium, or Low and provides a rationale.

## 2. Features

### 2.1. Chat Interface
- A responsive chat window to display conversations.
- A user input field for sending messages.
- Support for displaying messages from both the user and the assistant.
- Markdown rendering for chat messages.

### 2.2. Profile Configuration
- A dedicated profile screen to manage application settings.
- Ability to switch between the chat screen and the profile screen.

### 2.3. API Provider Management
- **Provider Selection:** Dropdown menu to select from a predefined list of API providers (Ollama, OpenAI, Gemini, Grok) or a custom provider.
- **Custom Base URL:** Input field for specifying a custom API base URL when the "Custom" provider is selected.
- **API Key Input:** Secure input field for entering the API key, with a button to toggle visibility.
- **Provider Verification:** Button to verify the provider and fetch available models.

### 2.4. Model Selection
- A dropdown menu to select the desired model from the list of models fetched from the verified provider.
- Button to save the selected provider and model settings.

### 2.5. System Prompt
- A Markdown editor (EasyMDE) for creating and editing the system prompt.
- Button to save the system prompt to local storage.
- A modal to prompt the user to set the system prompt if it's not already set.

### 2.6. PDF Upload
- A button to trigger a file input for selecting and uploading a PDF file.
- The application extracts the text from the uploaded PDF to be used in the chat conversation.

### 2.7. Chat History
- A modal to display the chat history in a sortable table.
- The table includes columns for Date, Model, API Provider, and Tokens.
- Ability to expand each row to view the full prompt and response.
- A button to clear the entire chat history.

### 2.8. Help Modal
- A modal that provides instructions on how to use the application.

### 2.9. Styling
- A modern user interface based on Material Design principles.
- Support for both light and dark color schemes, adapting to the user's system preference.

## 3. Technical Details

- **Frontend:** The application is built using HTML, CSS, and vanilla JavaScript.
- **Dependencies:** It utilizes external libraries such as:
    - `pdf.js` for PDF text extraction.
    - `easymde` for the Markdown editor.
    - `marked` and `showdown` for Markdown to HTML conversion.
- **Local Storage:** The application uses the browser's local storage to persist user settings, including:
    - API provider settings (including API keys).
    - The active provider.
    - The system prompt.
    - Chat history.

# Specification for RFP Eval HTML App

This document outlines the detailed specification for the "RFP Eval" application, a single-page web application designed to assist in evaluating Request for Proposals (RFPs).

## 1. Project Overview

RFP Eval is a client-side chat application that allows users to interact with a large language model (LLM) to evaluate RFPs. The user can configure the application to use various OpenAI-compatible APIs (like OpenAI, Gemini, or a local Ollama instance).

The core workflow involves:
1.  Configuring the API provider, model, and a "System Prompt".
2.  Uploading an RFP document (in PDF format).
3.  Asking the AI to evaluate the RFP based on the provided System Prompt.
4.  The AI grades the fitment as High, Medium, or Low and provides a rationale.

The application is built with HTML, CSS, and TypeScript/JavaScript, and it runs entirely in the browser, using local storage to persist user settings.

## 2. Application Architecture

-   **Client-Side SPA:** The application is a Single-Page Application (SPA) contained within a single `index.html` file.
-   **State Management:** Application state (API keys, selected model, System Prompt) is managed client-side using the browser's `localStorage`.
-   **API Interaction:** The application makes direct calls from the browser to the configured LLM API provider. There is no backend server component for the application itself.

## 3. File Structure

-   **/index.html**: The main entry point of the application. It contains the complete HTML structure for both the chat and profile screens, includes all necessary CSS and JavaScript libraries via CDN, and houses the core application logic within an inline `<script>` tag.
-   **/style.css**: Defines the application's visual style. It uses CSS variables to implement a Material Design-inspired theme with both light and dark modes. The layout is responsive, adapting to different screen sizes.
--  **/README.adoc**: The project's documentation file, containing setup instructions, API details, and a list of TODOs and known issues.

## 4. UI Components and Screens

The application consists of two main screens:

### 4.1. Chat Screen (`#chat-screen`)

This is the primary user interface for interacting with the AI.

-   **Header:** Displays the application title ("RFP Eval") and a "Profile" button to navigate to the settings screen.
-   **Chat Window (`#chat-window`):** Displays the conversation history between the user and the assistant. Messages are styled differently based on the sender.
-   **Chat Form (`#chat-form`):**
    -   **User Input:** A text field for typing messages.
    -   **File Upload Button (`#file-btn`):** A button that triggers a hidden file input to allow users to upload PDF files. Upon upload, the name of the file is displayed in the chat window.
    -   **Send Button:** Submits the message to the AI.

### 4.2. Profile Screen (`#profile-screen`)

This screen allows the user to configure the application's settings.

-   **Header:** Contains a "Back" button to return to the chat screen and the title "Profile".
-   **API Key Section:** An input field for the user to enter and save their API key.
-   **API Provider Section:**
    -   A dropdown (`#api-provider`) to select a provider: OpenAI, Gemini, Ollama, or Custom.
    -   A text input for the custom API base URL, which is shown only when "Custom" is selected.
    -   A "Verify Provider" button to fetch and populate the available models from the selected API endpoint.
-   **Model Selection Section:**
    -   A dropdown (`#model-select`) to choose from the models fetched from the provider.
    -   A "Save Model" button.
-   **System Prompt Section:**
    -   A rich text editor (`#company-profile-markdown`) powered by **EasyMDE** for inputting the System Prompt in Markdown format.
    -   A "Save Profile" button.

## 5. Core Functionality

### 5.1. API and Model Configuration

-   The user's API key, the chosen API provider's base URL, and the selected model ID are saved to `localStorage`.
-   When the user clicks "Verify Provider," the application sends a request to the `/v1/models` endpoint of the specified URL to fetch a list of available models.
-   This list populates the model selection dropdown.

### 5.2. System Prompt Management

-   The EasyMDE editor provides a user-friendly interface for writing the System Prompt in Markdown.
-   The profile content is saved to `localStorage` when the user clicks "Save Profile".

### 5.3. Chat and RFP Evaluation

-   Before sending a message, the application checks if the API key, provider URL, and model are configured. If not, it redirects the user to the profile screen.
-   **PDF Handling:** When a user uploads a PDF, the `pdf.js` library is used to extract the text content from the file. The name of the uploaded file is displayed in the chat window.
-   **Message Construction:** The extracted PDF text is prepended to the user's next message and sent to the AI as a single prompt. The system prompt instructs the AI to act as an RFP evaluation expert.
-   **API Call:** The application sends a `POST` request to the `/v1/chat/completions` endpoint of the configured API provider.
-   **Response Handling:** The AI's response is extracted from the JSON result. The `showdown.js` library is used to convert the Markdown-formatted response into HTML for display in the chat window.


## 6. External Dependencies

The application relies on the following third-party libraries, loaded via CDNs:

-   **pdf.js:** For client-side PDF parsing and text extraction.
-   **showdown.js & marked.js:** For converting Markdown text to HTML.
-   **EasyMDE:** To provide a rich Markdown editor for the System Prompt.

### 6.1. Note on Redundant Markdown Libraries

The application currently uses two different Markdown-to-HTML libraries: `showdown.js` and `marked.js`. This is not ideal, but it is necessary with the current implementation:
-   **`showdown.js`** is used by the core application logic to render chat messages from the AI.
-   **`marked.js`** is a required dependency for the **EasyMDE** editor, which uses it to generate the live preview in the "System Prompt" section.

Future work could involve refactoring the application to use a single Markdown library for both purposes.

## 7. Project Status (All Initial Requirements Completed)

All items from the initial "Known Issues and Future Work" list in the `README.adoc` have been addressed and are now considered complete.

-   **DONE:** Show timer updated every 3 seconds while awaiting LLM responce.
-   **DONE:** Log token consumption for each API request to a local database.
-   **DONE:** Implement a chat history feature, allowing users to view and clear past conversations, likely using a local database (e.g., IndexedDB).
-   **DONE:** Add a help icon to center top of the page, clicking on it will open a modal dialog with instructions on how to use the app.
-   **DONE:** add a button to show chat history, clicking on it will open a modal dialog with the chat history.
-   **DONE:** add a button to clear chat history, clicking on it will clear the chat history.
